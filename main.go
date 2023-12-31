package main

import (
	"api_go/routes"
	"fmt"
	"net/http"
)

const SERVER_PORT = "8000"

func main() {
	fmt.Println("Iniciando servidor na porta " + SERVER_PORT)
	routes.CarregaRotas()
	http.ListenAndServe(":"+SERVER_PORT, nil)
}
