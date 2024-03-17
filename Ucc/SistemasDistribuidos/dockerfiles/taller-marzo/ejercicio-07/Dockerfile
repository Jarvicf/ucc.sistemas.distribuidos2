# Especificar una imagen base
FROM python:3.8-slim

# Establecer el directorio de trabajo en la imagen
WORKDIR /app

# Copiar el archivo de requisitos de la aplicación
COPY requirements.txt .

# Instalar las dependencias de la aplicación
RUN pip install --no-cache-dir -r requirements.txt

# Copiar el código de la aplicación al directorio de trabajo en la imagen
COPY . .

# Exponer el puerto en el que la aplicación Flask estará escuchando
EXPOSE 5000

# Definir el comando por defecto para ejecutar la aplicación Flask
CMD ["flask", "run", "--host=0.0.0.0"]
