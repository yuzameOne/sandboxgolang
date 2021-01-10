package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO:
//  write file

var differenceBetweenIndex []int
var convertStringSliceToInt []int
var convertIntSliceToString []string
var finishStringWriteToFile string

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

			differenceBetweenIndex = append(differenceBetweenIndex, -(left - right))

		}

	}

	for _, val := range subStringLeft {

		value, _ := strconv.Atoi(val)

		convertStringSliceToInt = append(convertStringSliceToInt, value)
	}

	for i := convertStringSliceToInt[2]; i <= (convertStringSliceToInt[2] + differenceBetweenIndex[0]); i++ {
		convertStringSliceToInt[2] = i
		for j := convertStringSliceToInt[3]; j <= (convertStringSliceToInt[3] + differenceBetweenIndex[1]); j++ {

			convertStringSliceToInt[3] = j

			for _, val := range convertStringSliceToInt {

				value := strconv.Itoa(val)

				convertIntSliceToString = append(convertIntSliceToString, value)
			}

			fmt.Printf("type :%T , value : %s \n", convertIntSliceToString, convertIntSliceToString)

			finishStringWriteToFile = strings.Join(convertIntSliceToString, ".")

			fmt.Println(finishStringWriteToFile)

			for i := 0; i < len(convertIntSliceToString); {

				convertIntSliceToString[len(convertIntSliceToString)-1] = " "
				convertIntSliceToString = convertIntSliceToString[:len(convertIntSliceToString)-1]
				fmt.Println(convertIntSliceToString)
			}
			fmt.Printf("type :%T , value : %s \n", convertIntSliceToString, convertIntSliceToString)

			if j == differenceBetweenIndex[1] {
				convertStringSliceToInt[3] = 0
				break
			}

		}
	}

}

// file, err := os.OpenFile("finishIpDiap.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
// if err != nil {
// 	log.Fatalf("file not create : %s", err)
// }
