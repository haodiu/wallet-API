package main

import (
	"log"
	"wallet/routes"
)

func main() {
	r := routes.Engine()
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}