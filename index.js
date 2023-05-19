const express = require('express')
var cors = require('cors')

const app = express()
const port = 5000

app.use(cors())
app.use(express.json())

// Available routes 
app.use('/', require('./routes/Input'))

app.listen(port, () => {
  console.log(`Example app listening on port http://localhost:${port}`)
})