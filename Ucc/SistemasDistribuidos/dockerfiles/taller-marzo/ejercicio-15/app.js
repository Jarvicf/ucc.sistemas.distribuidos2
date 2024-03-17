// app.js
const express = require('express');
const app = express();

app.get('/', (req, res) => {
  res.send('¡Ejercicio 15 - Sistemas Distribuidos - Jarvic Balanta!');
});

app.listen(3000, () => {
  console.log('Servidor Express.js escuchando en el puerto 3000');
});
