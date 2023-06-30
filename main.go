package main

import (
	"log"
	"os"
	"website-analyser/app"
)

func main() {
	application := app.Generate()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	
	
}
