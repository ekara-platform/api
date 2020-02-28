package main

import (
	"flag"
	"github.com/ekara-platform/api/rest"
	"log"
	"os"
)

func main() {
	var logger = log.New(os.Stdout, "API  > ", log.LstdFlags)
	var port int
	flag.IntVar(&port, "port", 9999, "Port of the API")
	flag.Parse()

	logger.Printf("Ekara API on port %d...\n", port)
	rest.Init(logger, port)
	rest.Serve()
}
