# Utiliza la imagen base oficial de .NET Core SDK para compilar la aplicación
FROM mcr.microsoft.com/dotnet/sdk:latest AS build

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Restaura las dependencias y compila la aplicación
RUN dotnet restore
RUN dotnet publish -c Release -o out

# Utiliza la imagen base oficial de .NET Core para ejecutar la aplicación
FROM mcr.microsoft.com/dotnet/runtime:latest AS runtime

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos publicados de la etapa de compilación al directorio de trabajo
COPY --from=build /app/out .

# Comando para iniciar la aplicación
CMD ["dotnet", "MiApp.dll"]
