package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var nFlag = flag.Int("n", 1234, "help message for flag n")

	flag.Parse()

	log.Println(*nFlag)

	fmt.Println("cli")
}
