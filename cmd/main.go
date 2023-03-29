package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mariobac1/backend_webpages/infrastructure/handler"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/response"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	e := newHTTP(response.HTTPErrorHandler)

	dbPool, err := newDBConnection()
	fmt.Printf("Hay conexión %v", dbPool)
	fmt.Println("No hay conexión", err)
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(e, dbPool)

	port := os.Getenv("SERVER_PORT")
	if os.Getenv("IS_HTTPS") == "true" {
		err = e.StartTLS(":"+port, os.Getenv("CERT_PEM_FILE"), os.Getenv("KEY_PEM"))
	} else {
		err = e.Start(":" + os.Getenv("SERVER_PORT"))
	}
	if err != nil {
		log.Fatal(err)
	}

}
