package main

import (
	"fmt"
	"strings"
)

var finish = make([]string, 9)

func main() {

	startString := "80.92.160.0-80.92.175.255"

	// Split string like this [80 92 160 0-80 92 175 255]
	newString := strings.Split(startString, ".")

	//  take  index []string  0-80
	subStringMiddle := newString[3]

	// Split string like this []string  [0 80]
	newsubstring := strings.Split(subStringMiddle, "-")

	// get slice [80 92 160]
	leftSubstring := newString[0:3]

	// get slice [92 175 255]
	rigthSubstring := newString[4:]

	finish = append(leftSubstring, newsubstring...)
	finish = append(finish, rigthSubstring...)

	// result this [80 92 160 0 80 80 175 255]
	fmt.Println(finish)
	//  need result [80 92 160 0 80 92 175 255]
	fmt.Println(rigthSubstring)

}
