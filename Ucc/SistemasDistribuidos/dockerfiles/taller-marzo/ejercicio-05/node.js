const express = require('express');

const app = express();

app.get('/', (req, res) => {
    res.send("Ejercicio 5 - Sistemas Distribuidos - Jarvic Balanta")
})

app.listen(8080, () => {
    console.log("Escuchando por el puerto 8080")
})
