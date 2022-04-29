package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/router"
)

func main() {
	parameters := lib.Config()
	var port = "4000"

	if port = os.Getenv("PORT"); port == "" {
		port = parameters.PORT
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	originsOk := handlers.AllowedOrigins([]string{"*"})

	fmt.Printf("Listen and serve on port :%s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(headersOk, methodsOk, originsOk)(router.InitRoutes())); err != nil {
		fmt.Println(err)
	}
}
