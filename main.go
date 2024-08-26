package main

import (
	"log"
	"net/http"
	"os"

	"inventory-api/config"
	"inventory-api/routes"

	"github.com/sirupsen/logrus"
)

func main() {
    // Initialize logging
    file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    logrus.SetOutput(file)
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetLevel(logrus.InfoLevel)

    // Initialize configuration and database
    config.LoadConfig()
    config.ConnectDB()

    // Initialize routes
    r := routes.InitializeRoutes()

    // Start the server
    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}
