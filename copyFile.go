package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Incorrect amount of arguments...\n")
		os.Exit(1)
	}
	srcFile := args[0]
	destFile := args[1]
	fmt.Printf("Copying %s to %s\n", srcFile, destFile)

}

func usage() {
	fmt.Fprintf(os.Stderr, "help: ./copyFile src_file dest_file\n")
	flag.PrintDefaults()
	os.Exit(2)
}
