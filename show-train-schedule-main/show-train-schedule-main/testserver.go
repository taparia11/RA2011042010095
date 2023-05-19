package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	clients = make(map[uuid.UUID]RegisteredClient)
	quit    = make(chan bool)
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

type Train struct {
	Name           string     `json:"trainName"`
	Number         string     `json:"trainNumber"`
	DepartureTime  CustomTime `json:"departureTime"`
	SeatsAvailable Seat       `json:"seatsAvailable"`
	Price          Price      `json:"price"`
	DelayedBy      int        `json:"delayedBy"`
}

type CustomTime struct {
	Hours   uint8
	Minutes uint8
	Seconds uint8
}

func (c *CustomTime) String() string {
	return fmt.Sprintf("%v:%v:%v o'clock", c.Hours, c.Minutes, c.Seconds)
}

type Seat struct {
	Sleeper int `json:"sleeper"`
	AC      int `json:"AC"`
}

type Price struct {
	Sleeper int `json:"sleeper"`
	AC      int `json:"AC"`
}

// RegisteredClient ...
// Custom object which can be stored in the claims
type RegisteredClient struct {
	Name   string    `json:"companyName,omitempty"`
	ID     uuid.UUID `json:"clientID"`
	Secret string    `json:"clientSecret,omitempty"`
}

// AuthToken ...
// This is what is retured to the client
type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// AuthTokenClaim ...
// This is the cliam object which gets parsed from the authorization header
type AuthTokenClaim struct {
	*jwt.StandardClaims
	RegisteredClient
}

// ErrorMsg ...
// Custom error object
type ErrorMsg struct {
	Message string `json:"message"`
}

type Company struct {
	Name string `json:"companyName"`
}

func registerdHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var company Company
	_ = json.NewDecoder(req.Body).Decode(&company)
	if len(company.Name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMsg{Message: "company name is required"})
		return
	}
	id := uuid.New()
	client := RegisteredClient{
		Name:   company.Name,
		ID:     id,
		Secret: RandStringBytes(16),
	}
	clients[id] = client
	json.NewEncoder(w).Encode(client)
	return
}

func authHandler(w http.ResponseWriter, req *http.Request) {
	var client RegisteredClient
	_ = json.NewDecoder(req.Body).Decode(&client)

	expiresAt := time.Now().Add(time.Minute * 5).Unix()

	if len(client.ID) == 0 || len(client.Name) == 0 || len(client.Secret) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMsg{Message: "client ID, Name & Secret are mandatory fields"})
		return
	}

	if regClient, ok := clients[client.ID]; ok {
		if !reflect.DeepEqual(regClient, client) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorMsg{Message: "provided fields does not match with any of our registered client"})
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorMsg{Message: "given client is not registered"})
		return
	}
	claims := &AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		RegisteredClient{
			Name: client.Name,
			ID:   client.ID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(client.Secret))
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})
}

func validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		claims := &AuthTokenClaim{}
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					}

					if client, ok := clients[claims.ID]; ok {
						return []byte(client.Secret), nil
					}
					return nil, errors.New("not a registered client")
				})
				if err != nil {
					json.NewEncoder(w).Encode(ErrorMsg{Message: err.Error()})
					return
				}
				if token.Valid {
					registeredClient := claims.RegisteredClient

					if _, ok := clients[registeredClient.ID]; !ok {
						json.NewEncoder(w).Encode(ErrorMsg{Message: "invalid authorization token - does not contain reguistered Client ID"})
						return
					}
					context.Set(req, "decoded", claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token"})
				}
			} else {
				json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(ErrorMsg{Message: "An authorization header is required"})
		}
	})
}

func getTrainsScheduleHandler(w http.ResponseWriter, req *http.Request) {
	trains := make([]Train, 0)
	t := time.Now()
	for _, v := range schedules {
		if v.DepartureTime.Hours > uint8(t.Hour()) {
			trains = append(trains, v)
		} else if v.DepartureTime.Hours == uint8(t.Hour()) {
			if v.DepartureTime.Minutes > uint8(t.Minute()) {
				trains = append(trains, v)
			}
		}
	}

	json.NewEncoder(w).Encode(trains)
	return
}

func geTrainScheduleHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	trainNumber := vars["trainNumber"]
	if v, ok := schedules[trainNumber]; ok {
		json.NewEncoder(w).Encode(v)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func alterTrainParams() {
	for {
		select {
		case <-time.After(1 * time.Minute):
			n := rand.Intn(len(trainNumbers) - 1)
			tn := trainNumbers[n]
			train := schedules[tn]
			train.DelayedBy = n
			if n%2 == 0 {
				train.Price.AC += 10
				train.Price.Sleeper += 10
			} else {
				if train.Price.AC-15 > 0 {
					train.Price.AC -= 15
				}
				if train.Price.Sleeper-15 > 0 {
					train.Price.Sleeper -= 15
				}
			}
			if n%3 == 0 {
				if train.SeatsAvailable.AC-1 > 0 {
					train.SeatsAvailable.AC -= 1
				}
				if train.Price.Sleeper-2 > 0 {
					train.Price.Sleeper -= 2
				}
			}
			schedules[tn] = train
			fmt.Println("train params changed: ", schedules[tn])
		case <-quit:
			return
		}
	}
}

func main() {
	go alterTrainParams()
	router := mux.NewRouter()
	fmt.Println("Application Starting ...")
	router.HandleFunc("/register", registerdHandler).Methods("POST")
	router.HandleFunc("/auth", authHandler).Methods("POST")
	router.HandleFunc("/trains", validateTokenMiddleware(getTrainsScheduleHandler)).Methods("GET")
	router.HandleFunc("/trains/{trainNumber}", validateTokenMiddleware(geTrainScheduleHandler)).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
