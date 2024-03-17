# Utiliza la imagen base oficial de Ruby on Rails
FROM ruby:latest

# Establece el directorio de trabajo en /app
WORKDIR /app

# Instala las dependencias del sistema necesarias para la instalación de RubyGems
RUN apt-get update && apt-get install -y \
    build-essential \
    libpq-dev \
    nodejs \
 && rm -rf /var/lib/apt/lists/*

# Instala Bundler
RUN gem install bundler

# Copia los archivos de la aplicación al directorio de trabajo
COPY . .

# Instala las gemas del proyecto
RUN bundle install

# Ejecuta el servidor de Rails
CMD ["rails", "server", "-b", "0.0.0.0"]
