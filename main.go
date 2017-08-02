package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	var typeFlag InputType

	flag.Var(&typeFlag, "type", "Source type")
	flag.Parse()

	println(typeFlag)

	inputReader := bufio.NewScanner(os.Stdin)

	for inputReader.Scan() {
		path := inputReader.Text()
		log.Print(path)
	}
}
