package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO:
//  take for , and  twist   index  0 and 1
//  write file

var difference []int
var convertStringSliceToInt []int

func main() {
	startString := "80.92.160.0-80.92.175.255"

	newString := strings.Split(startString, "-")

	subStringLeft := strings.Split(string(newString[0]), ".")

	subStringRigth := strings.Split(string(newString[1]), ".")

	fmt.Sprintln(subStringRigth) // compiler scold, not use variable
	fmt.Sprintln(subStringLeft)  // compiler scold, not use variable

	for i := 0; i < len(subStringLeft); i++ {

		if subStringLeft[i] != subStringRigth[i] {

			left, _ := strconv.Atoi(subStringLeft[i])
			right, _ := strconv.Atoi(subStringRigth[i])

			difference = append(difference, -(left - right))

		}

	}

	for _, val := range subStringLeft {

		value, _ := strconv.Atoi(val)

		convertStringSliceToInt = append(convertStringSliceToInt, value)
	}

	for i := convertStringSliceToInt[2]; i <= (convertStringSliceToInt[2] + difference[0]); i++ {
		fmt.Println(i)
		for j := convertStringSliceToInt[3]; j <= (convertStringSliceToInt[3] + difference[1]); j++ {

			fmt.Println(j)
			if j > (j + 1) {
				break
			}
		}
	}

}
