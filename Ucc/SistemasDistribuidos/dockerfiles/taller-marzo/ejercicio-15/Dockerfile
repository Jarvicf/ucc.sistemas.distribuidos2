# Utiliza la imagen base oficial de Node.js
FROM node:latest

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Instala las dependencias del proyecto
RUN npm install

# Expone el puerto 3000 del contenedor
EXPOSE 3000

# Comando para iniciar el servidor Express.js
CMD ["node", "app.js"]
