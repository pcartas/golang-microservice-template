package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	r "github.com/pcartas/golang-lib/router"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := r.NewRouter(routes, generalMiddlewares)
	puerto := os.Getenv("PORT")

	fmt.Println("Microservicio Template en el puerto: " + puerto)
	server := http.ListenAndServe(":"+puerto, router)

	log.Fatal(server)

}
