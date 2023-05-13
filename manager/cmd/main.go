package main

import (
	"github.com/spf13/pflag"
	"log"
	"manager/internal/app"
)

var port = pflag.IntP("port", "p", 80, "")

func main() {
	pflag.Parse()

	if err := app.Run(*port); err != nil {
		log.Fatal(err)
	}
}
