from flask import Flask

# Crear una instancia de la aplicación Flask
app = Flask(__name__)

# Definir una ruta para el punto de entrada de la aplicación
@app.route('/')
def index():
    return '¡Hola. Este es el ejercicio 7 - Sistemas Distribuidos - Jarvic Balanta !'

# Punto de entrada para ejecutar la aplicación Flask
if __name__ == '__main__':
    app.run(debug=True)
