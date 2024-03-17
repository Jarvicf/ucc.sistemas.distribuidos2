# Utiliza la imagen base oficial de PostgreSQL
FROM postgres:latest

# Variables de entorno para la base de datos
ENV POSTGRES_DB=usuario
ENV POSTGRES_USER=java
ENV POSTGRES_PASSWORD=123456

# Copia el archivo de inicialización de la base de datos
COPY init.sql /docker-entrypoint-initdb.d/

