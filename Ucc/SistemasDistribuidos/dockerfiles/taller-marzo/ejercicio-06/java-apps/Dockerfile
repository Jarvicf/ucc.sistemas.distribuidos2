# Especificar una imagen base que incluya el JDK de Java
FROM openjdk:latest

# Establecer el directorio de trabajo en la imagen
WORKDIR /usr/src/app

# Copiar el código fuente de la aplicación al directorio de trabajo
COPY . .

# Compilar la aplicación
RUN javac Main.java

# Comando por defecto para ejecutar la aplicación
CMD ["java", "Main"]
