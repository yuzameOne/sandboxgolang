package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var newString strings.Builder

func main() {

	file, err := ioutil.ReadFile("unparseip.txt")

	if err != nil {
		log.Printf("file dont read : %s", err)
	}

}
