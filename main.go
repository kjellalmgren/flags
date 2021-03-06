package main

import (
	"flag"
	"flags/version"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	// TETRACON banner
	TETRACON = `
_________    __
|__    __|   | |
   |  |  ___ | |_   ____  ___   ___ ___  _ __ 
   |  | / _ \|  _| /  __|/ _ \ / __/ _ \| '_ \
   |  | \ __/| |_  | |  | (_| | (_| (_) | | | | 
   |__| \___| \__| |_|   \__,_|\___\___/|_| |_| 
   version: %s
   
   `
)

var (
	debug bool
	vrsn  bool
)

//
func init() {
	// parse flags
	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&vrsn, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.NewVersion()))
		flag.PrintDefaults()
	}

	flag.Parse()

	// set log level
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if vrsn {
		fmt.Printf("flag version %s\n", version.NewVersion())
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		return
	}

	// parse the arg
	arg := flag.Args()[0]

	if arg == "help" {
		usageAndExit("", 0)
	}

	if arg == "version" {
		fmt.Printf("flag version %s\n", version.NewVersion())
		os.Exit(0)
	}
}

// here we go
func main() {
	fmt.Println("Start...")
	fmt.Println("End...")
}

//
func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}
