const express = require('express');
const axios = require('axios')
const app = new express();

var lista = [];
var num = 0;
var h_redis = "http://104.154.91.245:8000"
var h_grpc = "http://35.232.241.29:8000"
var h_kafka = ""

function agregar(url, data) {
    axios
    .post(url, data)
    .then(res => {
      console.log("res")
    })
    .catch(error => {
      console.error("error")
    })   
}

app.use(express.json({ limit: '5mb', extended: true }));

app.post('/', async (req, res) => {
    const data = req.body;
    try {
        lista.push(data);
        num = Math.floor(Math.random()*(3-1))+1
        if(num == 1){
            agregar(h_grpc, data)
        }else if(num == 2){
            agregar(h_redis, data)
        }else{
            agregar(h_kafka, data)
        }
        res.json(data);
        console.log(num)
        
    } catch (err) {
        console.log(err);
        res.status(500).json({ 'message': 'failed' });
    }
});

app.get('/', (req, res) => {
    res.json(lista.pop())
});

app.listen(5000);
console.log("app running in port 5000")
