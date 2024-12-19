package main

import (
    "compartilhatech/internal/application/services"
    "compartilhatech/internal/infra/database/sqlc"
    "compartilhatech/internal/interface/api/controllers"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    dbConn := sqlc.NewDB()
    db, err := dbConn.Connect()
    if err != nil {
        log.Fatalln(err)
    }

    router := mux.NewRouter()

    personService := services.NewPersonService(db)

    controllers.NewPersonController(router, personService)

    corsHandler := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:5173"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    })

    handler := corsHandler.Handler(router)

    port := ":3333"
    fmt.Println("Starting server on port", port)
    if err := http.ListenAndServe(port, handler); err != nil {
        fmt.Println("Error:", err)
    }
}
