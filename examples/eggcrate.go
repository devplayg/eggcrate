package main

import (
	"fmt"
	"github.com/devplayg/eggcrate"
	"github.com/spf13/pflag"
	"os"
)

const (
	appName     = "EggCrate"
	description = "embeds static files into Go source"
)

var (
	fs     = pflag.NewFlagSet("EggCrate", pflag.ContinueOnError)
	dir    = fs.StringP("dir", "d", "", "Source directory")
	ext    = fs.StringP("ext", "e", "js,png,css,eot,svg,ttf,woff,woff2,js,html", "Comma separated extensions(ex: jpg,htm,...")
	output = fs.StringP("output", "o", "asset.go", "Output file")

	extMap = make(map[string]bool)
)

func init() {
	// Usage
	fs.Usage = func() {
		fmt.Printf("%s %s\n", appName, description)
		fs.PrintDefaults()
		os.Exit(1)
	}

	_ = fs.Parse(os.Args[1:])
}

func main() {
	if err := eggcrate.CheckDir(*dir); err != nil {
		fmt.Println(err)
		return
	}

	extensionMap, err := eggcrate.CreateExtensionMap(*ext)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := eggcrate.Encode(*dir, extensionMap, *output); err != nil {
		fmt.Println(err)
		return
	}
}
