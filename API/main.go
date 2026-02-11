package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gasparguilherme/Netwise/api/src/config"
	"github.com/gasparguilherme/Netwise/api/src/router"
)

func main() {
	config.Load()
	r := router.Generate()

	fmt.Printf("Escutando na porta %d", config.APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIPort), r))
}
