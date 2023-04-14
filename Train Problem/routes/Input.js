const express = require('express');
const router = express.Router();

// for register api
router.post('/register', async (req, res)=>{

    try {
    
    const{companyNam} = req.body;
   
   let data = await fetch("http://localhost:3000/register", {
        method: "POST",
        body: JSON.stringify({
          companyName : companyNam
        }),
        headers: {
          "Content-type": "application/json; charset=UTF-8"
        }
      });

    //   console.log(data)
      const dat = await  data.json()
    res.send(dat) 
        
} catch (error) {
    console.error(error.message);
    res.status(500).send("Internal Server Error");
}
})

//for Auth

router.post('/auth', async (req, res)=>{

    try {
    
    const{companyName,clientID,clientSecret} = req.body;
   let data = await fetch("http://localhost:3000/auth", {
        method: "POST",
        body: JSON.stringify({
            "companyName": companyName,
            "clientID": clientID,
            "clientSecret": clientSecret
        }),
        headers: {
          "Content-type": "application/json; charset=UTF-8"
        }
      });

    //   console.log(data)
      const dat = await  data.json()
    res.send(dat) 
        
} catch (error) {
    console.error(error.message);
    res.status(500).send("Internal Server Error");
}
})

// for List of trains
router.get('/trains', async (req, res)=>{
        try {
            const{authorization} = req.headers;
            console.log(authorization)
            let data = await fetch("http://localhost:3000/trains", {
                 method: "GET",
                 headers: {
                   "Content-type": "application/json; charset=UTF-8",
                   "Authorization" : "Bearer" + authorization
                 }
               });
         
             //   console.log(data)
               const dat = await  data.json()
             res.send(dat)   
        } catch (error) {
            console.error(error.message);
            res.status(500).send("Internal Server Error");
        }
    })

    // particular train number details
    router.get('/trains/:id', async (req, res)=>{
        try {
            const{authorization} = req.headers;
            const trainNo = req.params.id;
            let data = await fetch(`http://localhost:3000/trains/${trainNo}`, {
                 method: "GET",
                 headers: {
                   "Content-type": "application/json; charset=UTF-8",
                   "Authorization" : "Bearer" + authorization
                 }
               });
         
             //   console.log(data)
               const dat = await  data.json()
             res.send(dat)   
        } catch (error) {
            console.error(error.message);
            res.status(500).send("Internal Server Error");
        }
    })

module.exports = router