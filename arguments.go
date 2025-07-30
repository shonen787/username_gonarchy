package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	input_filename     string
	help, list_formats bool
}

var options Options

func init() {
	flag.StringVar(&options.input_filename, "input", "", "Input Filename")
	flag.StringVar(&options.input_filename, "i", "", "Input Filename (shorthand)")
	flag.BoolVar(&options.help, "help", false, "Help Message")
	flag.BoolVar(&options.help, "h", false, "Help Message (Shorthand)")
	flag.BoolVar(&options.list_formats, "list", false, "List Morphs")
	flag.BoolVar(&options.list_formats, "l", false, "List Morphs (Shorthand)")
	flag.Parse()

	if options.help {
		flag.Usage()
		os.Exit(0)
	}

	if options.list_formats {
		for _, morph := range morphs {
			fmt.Println("Morph: ", morph.Name())
		}
		os.Exit(0)
	}

	if options.input_filename == "" {
		fmt.Println("[!] Input necessary")
		flag.PrintDefaults()
		os.Exit(0)
	}

}
