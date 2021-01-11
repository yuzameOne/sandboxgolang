package main

import (
	"log"
	"os"
)

func read() {

	file, err := os.Open("unparseip.txt")

	if err != nil {
		log.Fatalf("file not open", err)
	}

	defer file.Close()
}
