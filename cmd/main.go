package main

import (
	"log"
	"os"

	"github.com/mariobac1/backend_webpages/domain/login"
	"github.com/mariobac1/backend_webpages/infrastructure/handler"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/response"
)

func main() {

	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = login.LoadFiles(os.Getenv("PRIVATE_RSA"), os.Getenv("PUBLIC_RSA"))
	if err != nil {
		log.Fatalf("No se pudo cargar los certificates: %v", err)
	}

	err = validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	e := newHTTP(response.HTTPErrorHandler)

	dbPool, err := newDBConnection()
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
