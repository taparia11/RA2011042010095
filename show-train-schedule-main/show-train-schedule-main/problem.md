# Trains Schedule Display

## Create a real-time Train Schedule Display of all trains in the next 12 hours along with the seat availability and pricing.

We would like to offer below services to our customers:
- Display real-time train schedules of all trains along with seat availability, prices
- Trains departing in the next 30 minutes should be ignored.
- Trains should be displayed in the ascending order of price, descending order of tickets and descending order of departure time(after considering delays)
- The response given to the user has to be ordered based on the current values of the above three values.
- Trains that fulfil the allowed time window after delays should also be considered
- Each train contains seat availability and prices for 2 train coach types - sleeper and AC
- The prices, seatsAvailability of all tickets are subject to change based on market conditions such as demand, supply, departure time. The latest price only has to be considered
- API calls of the John Doe Railway Server are chargeable (Number of calls has to be minimised)
- API server built by you has to be performant providing users a timely and effective response in the shortest time possible

__Create an API server that follows REST Standard to display the real-time Schedule based on the criteria mentioned__

## John Doe Railway Server APIs

### /register (POST)

This is an API to register your company with the John Doe Railway Server.

#### Request

```
{
    "companyName": "something"
}
```

#### Response Expected: (Status Code: 200)

```
{
    "companyName": "Affordmed",
    "clientID": "59e73b59-f659-4b08-892d-c96342bcf225",
    "clientSecret": "atwtPilfsfykSBGp"
}
```

### /trains (GET)

This is an API which returns all the train details from the John Doe Railway Server. This is a protected route and requires you to provide the Authorisation Token in the Header.

#### Response Expected: (Status Code: 200)

```
[
    {
        "trainName": "somthing exp",
        "trainNumber": "1234",
        "departureTime": 
        {
            "Hours": 21,
            "Minutes": 35,
            "Seconds": 0
        },
        "seatAvailable": 
        {
            "sleeper": 5,
            "AC": 2
        },
        "price": 
        {
            "sleeper": 432,
            "AC": 1232
        },
        "delayedBy": 10 // in minutes
    },
    {
        "trainName": "exp",
        "trainNumber": "122",
        "departureTime": 
        {
            "Hours": 21,
            "Minutes": 35,
            "Seconds": 0
        },
        "seatsAvailable": 
        {
            "sleeper": 5,
            "AC": 2
        },
        "price": 
        {
            "sleeper": 432,
            "AC": 1232
        },
        "delayedBy": 10 // in minutes
    }
]
```

### /trains/{trainNumber} (GET)

This is an API which returns the details of a particular train from the John Doe Railway Server. This is a protected route and requires you to provide the Authorisation Token in the Header.

#### Response Expected: (Status Code: 200)

```
{
    "trainName": "somthing exp",
    "trainNumber": "1234",
    "departureTime": 
    {
        "Hours": 21,
        "Minutes": 35,
        "Seconds": 0
    },
    "seatsAvailable": 
    {
        "sleeper": 5,
        "AC": 2
    },
    "price": 
    {
        "sleeper": 432,
        "AC": 1232
    },
    "delayedBy": 10 // in minutes
}
```

### /auth (POST)

This is an API to obtain the Authorisation Token for your company from the John Doe Railway Server.

#### Request

```
{
    "companyName": "Affordmed",
    "clientID": "0fc91eb2-5008-4eb5-b312-ba6dc8f65e5b",
    "clientSecret": "oJnNPGsiuzytMOJP"
}
```

#### Response Expected: (Status Code: 200)

```
{
    "token_type": "Bearer",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODE0NDk5MDksImNvbXBhbnlOYW1lIjoiQWZmb3JkbWVkIiwiY2xpZW50SUQiOiIwZmM5MWViMi01MDA4LTRlYjUtYjMxMi1iYTZkYzhmNjVlNWIifQ.5rrWy0NzpHhrOBzR1KHhAai0HxBOWe7gQ89IvIjoABA",
    "expires_in": 1681449909
}
```
