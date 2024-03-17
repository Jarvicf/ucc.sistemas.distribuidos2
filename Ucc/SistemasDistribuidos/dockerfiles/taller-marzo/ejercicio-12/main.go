package main

import (
	"log"
	"net/http"
)

func main() {
	// Configuración del enrutador HTTP
	mux := http.NewServeMux()

	// Manejador para la ruta raíz
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Devuelve una página HTML simple que incluye una imagen y un mensaje
		html := `
			<html>
				<head>
					<title>ejercicio12</title>
				</head>
				<body>
					<h1>Hola</h1>
					<p>Este es el ejercicio 12 - Sistemas Distribuidos UCC</p>
					<img src="/image/ucc.png" alt="Gopher" width="400" height="200">
				</body>
			</html>
		`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	// Configuración del servidor HTTP para servir la imagen
	mux.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("."))))

	// Configuración del servidor HTTP
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Inicia el servidor
	log.Println("Servidor escuchando en el puerto 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %s", err)
	}
}
