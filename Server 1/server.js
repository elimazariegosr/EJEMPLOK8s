const express = require('express');
const http = require('http');

const app = new express();
app.use(express.json({ extended: true }))

app.post('/', (req, res) => {
    var data = req.body;
    data['first'] = 'Server 1' 
    const newData = JSON.stringify(data);
    const options = {
        hostname: '35.226.102.166',
        port: 80,
        path: '/',
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Content-Length': Buffer.byteLength(newData)
        },
    }
    const request = http.request(options, (response) => {
        response.setEncoding('utf8');
        response.on('data', (d) => {
            res.json(JSON.parse(d.toString()));
        });
    });
    request.on('error', (error) => {
        res.status(500).json(error.toString());
    });
    request.write(newData);
    request.end();
});

app.get('/', (req, res) => {
    res.json({message: 'OK'})
});

app.listen(3000);
console.log("corriendo en el 3k");