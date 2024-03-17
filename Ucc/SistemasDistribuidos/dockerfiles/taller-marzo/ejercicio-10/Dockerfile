# Utiliza la imagen base oficial de PHP con Apache
FROM php:7.4-apache

# Copia los archivos de tu aplicación PHP al directorio de trabajo del servidor Apache
COPY . /var/www/html

# Expone el puerto 80 del contenedor
EXPOSE 80

# Establece el ServerName para suprimir la advertencia
RUN echo "ServerName localhost" >> /etc/apache2/apache2.conf

# Reinicia Apache para aplicar los cambios
RUN service apache2 restart

# Comando para iniciar el servidor Apache en primer plano
CMD ["apache2-foreground"]
