package main

import (
	"fmt"

	"github.com/Jose-Salazar-27/prueba_tecnica/server"
)

func main() {
	fmt.Println("started")

	server := server.NewServer(":9000")
	server.Init()

}
