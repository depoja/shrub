package main

import (
	"flag"
	"os"
	"shrub/file"
	"shrub/repl"
)

func main() {
	// Handle program flags
	isRepl := flag.Bool("repl", false, "run the shrub REPL")
	flag.Parse()

	if *isRepl {
		repl.Repl()
	} else {
		// Handle program arguments
		if len(os.Args) < 2 {
			panic("No file provided!")
		}

		file.Interpret(os.Args[1])
	}
}
