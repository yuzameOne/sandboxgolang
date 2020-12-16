package main

import (
	"fmt"
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

	for _, v := range file {

		newString.WriteByte(v)

		// 10 == "\n"
		if v == 10 {

			fmt.Println(newString.String())

			newString.Reset()

		}
	}
}
