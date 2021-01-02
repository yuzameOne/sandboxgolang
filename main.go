package main

import (
	"fmt"
	"strconv"
	"strings"
)

var newFinishSlice []string
var convertStringSliceToInt = make([]int, len(newFinishSlice))

func main() {
	startString := "80.92.160.0-80.92.175.255"

	newString := strings.Split(startString, "-")

	subStringLeft := strings.Split(string(newString[0]), ".")

	subStringRigth := strings.Split(string(newString[1]), ".")

	newFinishSlice = append(subStringLeft, subStringRigth...)

	fmt.Printf("type slice : %T , %v \n", newFinishSlice, newFinishSlice)

	for _, val := range newFinishSlice {

		value, _ := strconv.Atoi(val)

		convertStringSliceToInt = append(convertStringSliceToInt, value)
	}

	fmt.Printf("type slice : %T , %v \n", convertStringSliceToInt, convertStringSliceToInt)

	for i := 0; i < 4; i++ {
		fmt.Println(convertStringSliceToInt[i])
		for j := i + 4; j < 8; j++ {
			fmt.Println(convertStringSliceToInt[j])
			if i != j {
				break
			}
		}
	}

}
