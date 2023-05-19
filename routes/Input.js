const express = require('express');
const router = express.Router();

// get method bfhl
router.get('/bfhl', async (req, res)=>{
        try {
            let data = {
              "opration_code":1
            }
             res.status(200).json(data)   
        } catch (error) {
            console.error(error.message);
            res.status(500).send("Internal Server Error");
        }
    })

    

    // post method bfhl
    router.post('/bfhl', async (req, res)=>{
        try {
            let {data}  = req.body
            console.log(data)
            
            // let length = data.length;
            let string_data = []
            let number_data = []
            for (let index = 0; index < data.length; index++) {
              if (typeof data[index] == "number") {
                number_data.push(data[index].toString())
              }
              else{
                string_data.push(data[index])
              }
            }
            var response = {
              "is_success": true,
              "user_id": "nikhil_taparia_17092002", 
              "email" : "nt8770@srmist.edu.in",
              "roll_number":"RA2011042010095",
              "numbers": number_data,
              "alphabets": string_data
            }

            res.status(200).json(response)
        } catch (error) {
            console.error(error.message);
            res.status(500).send("Internal Server Error");
        }
    })

module.exports = router