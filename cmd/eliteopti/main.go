// cmd/eliteopti/main.go
package main

import (
	"flag"
	"log"
	"os"

	"eliteopti/internal/eliteopti"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := eliteopti.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
