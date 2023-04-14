const express = require('express');
const router = express.Router();


router.get('/numbers', async (req, res)=>{
        try {
          const url = req.query.url;
          // console.log(url[0])
          let number = []


          for( let count = 0; count < url.length; count++) {
            let lent = url[count].length
            let data = await fetch(`http://localhost:8090/${url[count].substr(22,lent)}`, {
                 method: "GET",
                 headers: {
                   "Content-type": "application/json; charset=UTF-8",
                 }
               });
              const dat = await  data.json()
              const combined1 = number.concat(dat.numbers);
              number = combined1
              // console.log(number)
              }

              let unique = [...new Set(number)];
             res.send(unique.sort(function(a, b){return a - b}))   
        } catch (error) {
            console.error(error.message);
            res.status(500).send("Internal Server Error");
        }
    })


module.exports = router