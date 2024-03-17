# Utiliza la imagen base oficial de Python
FROM python:latest

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Instala las dependencias del proyecto
RUN pip install -r requirements.txt

# Ejecuta las migraciones de la base de datos (si es necesario)
RUN python manage.py migrate

# Recolecta los archivos estáticos (si es necesario)
RUN python manage.py collectstatic --noinput

# Expone el puerto 8000 del contenedor
EXPOSE 8000

# Comando para iniciar el servidor Django
CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]
