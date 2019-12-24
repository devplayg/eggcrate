package main

import (
	"errors"
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
	fs         = pflag.NewFlagSet("EggCrate", pflag.ContinueOnError)
	dir        = fs.StringP("dir", "d", "", "Source directory")
	extensions = fs.StringP("ext", "e", "js,png,css,eot,svg,ttf,woff,woff2,js,html", "Comma separated extensions(ex: jpg,htm,...")
	output     = fs.StringP("output", "o", "asset.go", "Output file")

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
	if err := CheckDir(*dir); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := eggcrate.Encode(*dir, *extensions, *output); err != nil {
		fmt.Println(err)
		return
	}
}

func CheckDir(path string) error {
	if len(path) < 1 {
		return errors.New("empty directory")
	}

	dir, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New("directory not found: " + path)
	}

	if !dir.IsDir() {
		return errors.New("invalid directory: " + path)
	}
	return nil
}
