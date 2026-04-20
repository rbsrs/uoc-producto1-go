package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Puerto configurable por variable de entorno (útil en Docker)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// 1) Servir carpeta /static/ (imagen)
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2) Endpoint principal
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Buenas prácticas: método permitido y content-type
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprint(w, `


  
  
  


  
Soy alumno de la UOC

  

Producto 1: entorno Go + Docker

  UOC

`)
	})

	// 3) Endpoint extra opcional: salud
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, "OK")
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: logRequest(mux),
	}

	log.Printf("Servidor escuchando en http://localhost:%s", port)
	log.Fatal(srv.ListenAndServe())
}

// Middleware simple de logs
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
