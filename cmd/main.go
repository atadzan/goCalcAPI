package main

import (
	"flag"
	"log"

	"github.com/atadzan/goCalcAPI/pkg/handlers"
	"github.com/atadzan/goCalcAPI/pkg/service"
	"github.com/atadzan/goCalcAPI/server"
)

func main() {
	port := flag.String("port", "8080", "listening port of http server")
	flag.Parse()
	serviceInstance := service.New()
	handlerInstance := handlers.New(serviceInstance)

	if err := server.Run(handlerInstance, *port); err != nil {
		log.Fatalln("can't run server. Err: ", err)
	}
}
