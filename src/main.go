package main

import (
	"log"
	"net/http"

	"anthonypillot.com/starter-golang/controller"
	"anthonypillot.com/starter-golang/dao"
	"anthonypillot.com/starter-golang/service"
)

func main() {
	// Define all DAO
	photoDAO := dao.NewPhotoDAO("https://jsonplaceholder.typicode.com")

	// Define all services
	photoService := service.NewPhotoService(*photoDAO)

	// Define all controllers
	controller.NewPhotoController(*photoService)

	// Define all middlewares
	loggedMux := loggingMiddleware(http.DefaultServeMux)

	port := ":8080"
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(port, loggedMux))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Received request: %s %s", request.Method, request.URL.Path)
		next.ServeHTTP(writer, request)
	})
}
