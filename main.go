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

	// file, err := os.OpenFile("finishIpDiap.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	// if err != nil {
	// 	log.Fatalf("file not create : %s", err)
	// }

	for i := convertStringSliceToInt[2]; i <= (convertStringSliceToInt[2] + difference[0]); i++ {
		convertStringSliceToInt[2] = i
		fmt.Println(convertStringSliceToInt)
		for j := convertStringSliceToInt[3]; j <= (convertStringSliceToInt[3] + difference[1]); j++ {

			convertStringSliceToInt[3] = j
			fmt.Println(convertStringSliceToInt)
			if j == difference[1] {
				convertStringSliceToInt[3] = 0
				break
			}
		}
	}

}
