package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// TODO:
// read uparseip.txt 
//  swap variable startString in loop 

var differenceBetweenIndex []int
var convertStringSliceToInt []int
var convertIntSliceToString []string
var finishStringWriteToFile string

func main() {
	startString := "80.92.160.0-80.92.175.255"

	newString := strings.Split(startString, "-")

	subStringLeft := strings.Split(string(newString[0]), ".")

	subStringRigth := strings.Split(string(newString[1]), ".")

	twoIndexStaticNumber, _ := strconv.Atoi(subStringLeft[2])

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

		if i == (twoIndexStaticNumber+differenceBetweenIndex[0])+1 {
			break
		}
		convertStringSliceToInt[2] = i

		for j := convertStringSliceToInt[3]; j <= (convertStringSliceToInt[3] + differenceBetweenIndex[1]); j++ {

			convertStringSliceToInt[3] = j

			for _, val := range convertStringSliceToInt {

				value := strconv.Itoa(val)

				convertIntSliceToString = append(convertIntSliceToString, value)
			}

			finishStringWriteToFile = strings.Join(convertIntSliceToString, ".")

			for i := 0; i < len(convertIntSliceToString); {

				convertIntSliceToString[len(convertIntSliceToString)-1] = " "
				convertIntSliceToString = convertIntSliceToString[:len(convertIntSliceToString)-1]
			}

			file, err := os.OpenFile("finishIpDiap.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
			if err != nil {
				log.Fatalf("file not create : %s", err)
			}

			file.WriteString(finishStringWriteToFile + "\n")

			if j == differenceBetweenIndex[1] {
				convertStringSliceToInt[3] = 0
				break
			}

		}

	}

}
