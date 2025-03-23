package main

import (
	"log"
	"os"

	"github.com/ninosistemas10/delivery/infrastructure/handler"
	"github.com/ninosistemas10/delivery/infrastructure/handler/response"
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
	e.Static("/promocion", "uploads/promocion")
	e.Static("/categorias", "uploads/categorias")

	dbPool, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(e, dbPool)

	err = e.Start(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}

}
