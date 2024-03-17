# Utilizar la imagen base oficial de Nginx
FROM nginx:latest

# Copiar el archivo de configuración personalizado de Nginx al directorio de configuración
COPY nginx.conf /etc/nginx/nginx.conf

# Exponer el puerto en el que Nginx estará escuchando
EXPOSE 4000

# Comando por defecto para ejecutar Nginx
CMD ["nginx", "-g", "daemon off;"]
