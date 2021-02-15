package main

import (
	"log"

	"github.com/1k-ct/indecisive/router"
)

func main() {
	r := router.InitRouters()
	if err := r.Run(); err != nil {
		log.Fatalf("server error: %s", err)
	}
}
