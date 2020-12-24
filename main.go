package main

import (
	"fmt"
	"strings"
)

var stringLength int
var finishSlice []string

func main() {

	s := "80.92.160.0-80.92.175.255"

	// find full  string length

	// for i := 0; i < len(s); i++ {

	// 	stringLength = i
	// 	// fmt.Println(stringLength)
	// }

	newString := strings.Split(s, ".")

	substring := newString[3]
	substringleft := newString[3:]

	newsubstring := strings.Split(substring, "-")

	finishSlice := append(finishSlice, newString, substringleft)

	// TODO variable newsubstring append in newString
	// How copy slice 2 part??????

	fmt.Printf("basic []string : %v \n", newString)
	fmt.Printf("substring  type string : %v \n", substring)
	fmt.Printf("copy index value  []string : %v \n", newsubstring)

	fmt.Printf("finish  []string : %v ", finishSlice)

	// fmt.Println(newString[3])
	// fmt.Println(newString[6])

	// newNewString = append(newNewString, )

}
