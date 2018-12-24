package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("test")
	log.SetOutput(os.Stdout)

}
