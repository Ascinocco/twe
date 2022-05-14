package main

import (
	"TheWarEconomy/api/controllers"
	"TheWarEconomy/api/middleware"
	"TheWarEconomy/api/utils"
	"fmt"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		fmt.Println("Couldn't load env values", envErr)
		os.Exit(1)
	}

	port := os.Getenv(utils.EnvPort)

	router := mux.NewRouter()
	router.Use(middleware.JwtVerification)

	router.HandleFunc("/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/session/create", controllers.CreateSession).Methods("POST")
	router.HandleFunc("/pmc/create", controllers.CreatePmc).Methods("POST")
	router.HandleFunc("/ws", controllers.EstablishWsConn)
	router.Handle("/", router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	err := http.ListenAndServe(":"+port, handler)

	if err != nil {
		fmt.Println("Error booting up http server", err)
		os.Exit(1)
	}
}
