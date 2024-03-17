# Utiliza la imagen base oficial de Go
FROM golang:latest

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Compila la aplicación
RUN go build -o myapp .

# Expone el puerto 8080 del contenedor
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./myapp"]
