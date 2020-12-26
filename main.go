package main

import (
	"fmt"
	"strings"
)

var newFinish []string

func main() {
	startString := "80.92.160.0-80.92.175.255"

	newString := strings.Split(startString, "-")

	subStringLeft := strings.Split(string(newString[0]), ".")

	subStringRigth := strings.Split(string(newString[1]), ".")

	newFinish = append(subStringLeft, subStringRigth...)

	fmt.Println(newFinish)

}
