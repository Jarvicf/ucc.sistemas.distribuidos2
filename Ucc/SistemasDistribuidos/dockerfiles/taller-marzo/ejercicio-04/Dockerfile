# Utilizamos la imagen oficial de MySQL
FROM mysql:latest

#Credenciales
ENV MYSQL_ROOT_PASSWORD=123456
ENV MYSQL_DATABASE=usuarios
ENV MYSQL_USER=java
ENV MYSQL_PASSWORD=123456

COPY init.sql /docker-entrypoint-initdb.d/
EXPOSE 3306
