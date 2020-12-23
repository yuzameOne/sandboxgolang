package main

import (
	"fmt"
	"strings"
)

var stringLength int

func main() {

	s := "80.92.160.0-80.92.175.255"

	// find full  string length

	// for i := 0; i < len(s); i++ {

	// 	stringLength = i
	// 	// fmt.Println(stringLength)
	// }

	newString := strings.Split(s, ".")

	substring := newString[3]

	newsubstring := strings.Split(substring, "-")

	// TODO variable newsubstring append in newString
	// How copy slice 2 part??????

	fmt.Println(substring)
	fmt.Println(newsubstring)
	fmt.Println(newString)

	fmt.Println(newString[3])
	fmt.Println(newString[6])

}
