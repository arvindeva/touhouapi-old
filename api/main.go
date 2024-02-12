package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/arvindeva/touhouapi/api/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port, exists := os.LookupEnv("PORT")
	if exists {
		fmt.Println(port)
	} else {
		return
	}

	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!")
	})
	r.HandleFunc("/touhou", handlers.GetTouhous)
	r.HandleFunc("/touhou/{id}", handlers.GetTouhouById)

	http.ListenAndServe(":"+port, r)
}
