package main

import (
	"fmt"
	"github.com/devplayg/eggcrate"
)

func main() {
	config := eggcrate.Config{
		Dir:        "/tmp/static",
		OutFile:    "output.go",
		UriPrefix:  "/assets",
		Extensions: "js,css",
	}
	_, err := eggcrate.Encode(&config)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
