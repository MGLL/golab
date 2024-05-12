package main

import (
	_ "github.com/mgll/golab/simple-http-server/entrypoints"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	runServer()
}

func runServer() {
	log.Println("Starting http server...")
	log.Printf("Started on port %s", port)
	err := http.ListenAndServe(port, nil)
	{
		if err != nil {
			log.Fatal(err)
		}
	}
}
