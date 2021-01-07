package main

import (
	"fmt"
	"strconv"
	"strings"
)

// var newFinishSlice []string
// var convertStringSliceToInt = make([]int, len(newFinishSlice))
// var thirdNumber int
// var fourNumber int

// TODO:
//  create  []int slice, append to []int slice varable difference  index  [0,1]
// 	convert  subStringLeft   to int[] slice
//  take for , and  twist   index  0 and 1
//  write file

func main() {
	startString := "80.92.160.0-80.92.175.255"

	newString := strings.Split(startString, "-")

	subStringLeft := strings.Split(string(newString[0]), ".")

	subStringRigth := strings.Split(string(newString[1]), ".")

	fmt.Sprintln(subStringRigth) // compiler scold, not use variable
	fmt.Sprintln(subStringLeft)  // compiler scold, not use variable

	// fmt.Printf("type  subStringLeft : %T \n", subStringLeft)

	for i := 0; i < len(subStringLeft); i++ {

		if subStringLeft[i] != subStringRigth[i] {

			fmt.Printf("left : %s ,  right : %s\n", subStringLeft[i], subStringRigth[i])

			left, _ := strconv.Atoi(subStringLeft[i])
			right, _ := strconv.Atoi(subStringRigth[i])

			difference := left - right

			fmt.Printf(" type : %T,  %d \n", -difference, -difference)
		}

		// newFinishSlice = append(subStringLeft, subStringRigth...)

		// fmt.Printf("type slice : %T , %v \n", newFinishSlice, newFinishSlice)

		// for _, val := range newFinishSlice {

		// 	value, _ := strconv.Atoi(val)

		// 	convertStringSliceToInt = append(convertStringSliceToInt, value)
		// }

		// fmt.Printf("type slice : %T , %v \n", convertStringSliceToInt, convertStringSliceToInt)

		// for i := 0; i < 4; i++ {
		// 	fmt.Println(convertStringSliceToInt[i])
		// 	for j := i + 4; j < 8; j++ {
		// 		fmt.Println(convertStringSliceToInt[j])
		// 		if i != j && i < j {
		// 			if i < j {
		// 				thirdNumber = convertStringSliceToInt[j] - convertStringSliceToInt[i]
		// 				fmt.Println(thirdNumber)
		// 			}
		// 			break
		// 		}
		// 	}
		// }
	}
}
