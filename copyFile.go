package main

import (
	"flag"
	"fmt"
	"os"
    "io/ioutil"
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
    writeFile(destFile, readFile(srcFile))
    fmt.Print("File written! Checking....\n")
    writtenData := readFile(destFile)
    fmt.Printf("this is in the destination file:\n%s\n", writtenData)

}

func checkErr(e error){
    if e != nil{
        panic(e)
    }
}

func readFile(sourceFile string) []byte {
   data, err := ioutil.ReadFile(sourceFile)
   checkErr(err)

   fmt.Printf("Here's what I read from %s:\n%s\n", sourceFile, data)
   return data
}

func writeFile(destFile string,data []byte) bool {
   err := ioutil.WriteFile(destFile, data, 0644)
   defer checkErr(err)
   return err != nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "help: ./copyFile src_file dest_file\n")
	flag.PrintDefaults()
	os.Exit(2)
}
