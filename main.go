package main

import (
	"fmt"
	"strings"
)

var finish []string

func main() {

	startString := "80.92.160.0-80.92.175.255"

	// Split string like this [80 92 160 0-80 92 175 255]
	newString := strings.Split(startString, ".")

	//  take  index []string  0-80
	subStringMiddle := newString[3]

	// Split string like this []string  [0 80]
	newsubstring := strings.Split(subStringMiddle, "-")

	// TODO merge 3 slice in one slice

	leftSubstring := newString[0:3]
	rigthSubstring := newString[4:]

	fmt.Println(newsubstring)

	finish = append(finish, leftSubstring, rigthSubstring...)

	fmt.Println(finish)

}
