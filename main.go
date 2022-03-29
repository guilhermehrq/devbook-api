package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"net/http"
)

func main() {
	config.LoadEnv()

	r := router.Generate()

	fmt.Printf("API listening on the port %d!\n", config.APIPort)
	http.ListenAndServe(fmt.Sprintf(":%d", config.APIPort), r)
}
