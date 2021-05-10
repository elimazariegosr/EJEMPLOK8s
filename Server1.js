const express = require('express')
var bodyParser = require('body-parser');
const cors = require('cors');

const app = express()

const port = 3000
var corsOptions = { origin: true, optionsSuccessStatus: 200 };
app.use(cors(corsOptions));
app.use(bodyParser.json({ limit: '10mb', extended: true }));
app.use(bodyParser.urlencoded({ limit: '10mb', extended: true }))

app.get('/', (req, res) => {
  res.send('Hello World!')
})

app.post('/eje1',(req, res)=>{
    let body = req.body;
    let value = body.value;
    console.log("server 1")
    res.send(value)
})
app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
