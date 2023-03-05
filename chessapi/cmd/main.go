package main

import (
	"log"

	"github.com/dchudik/chessapp/chessapi"
)

func main() {
	err := chessapi.RunHttpServer()
	log.Fatal(err)
}
