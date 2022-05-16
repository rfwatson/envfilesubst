package main

import (
	"flag"
	"fmt"
	"os"

	"git.netflux.io/rob/envfilesubst/scanner"
)

const version = "0.1"

func main() {
	var (
		path         string
		printVersion bool
	)

	flag.StringVar(&path, "f", "", "The envfile to read from")
	flag.BoolVar(&printVersion, "version", false, "Print version and exit")
	flag.Parse()

	if printVersion {
		fmt.Fprintf(os.Stderr, "envfilesubst version %s\n", version)
		os.Exit(0)
	}

	if path == "" {
		fmt.Fprint(os.Stderr, "envfilesubst reads an input from stdin and pipes it to stdout, replacing $ENV_VAR or ${ENV_VAR} occurrences that can be found in the provided envfile.\n\n")
		fmt.Fprint(os.Stderr, "Variables that cannot be found in envfile are left unchanged.\n\n")
		fmt.Fprint(os.Stderr, "Usage:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	envfile, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	scanner := scanner.New(os.Stdout, os.Stdin, envfile)
	if err := scanner.Scan(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
