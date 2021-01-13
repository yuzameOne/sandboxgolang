package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// TODO:
// remove element  in loop  differenceBetweenIndex []int ( added like this [0 255 0 15])
// remove element  in loop  convertStringSliceToInt []int (added like this [5 100 67 0 82 162 35 240]  )

var differenceBetweenIndex []int
var convertStringSliceToInt []int
var convertIntSliceToString []string
var finishStringWriteToFile string
var parseIPLinesFile []string

func main() {

	//  open file
	file, err := os.Open("unparseip.txt")

	// check error
	if err != nil {
		log.Fatal("file not Open", err)
	}
	//  read file inside all bytes
	scanner := bufio.NewScanner(file)

	// read file before \n
	scanner.Split(bufio.ScanLines)

	// write newline string in []string
	for scanner.Scan() {
		parseIPLinesFile = append(parseIPLinesFile, scanner.Text())
	}

	//  read []string  in loop
	for i := 0; i < len(parseIPLinesFile); i++ {

		// take string like  "5.100.67.0-5.100.67.255"
		startString := parseIPLinesFile[i]
		//  split string in []string
		newString := strings.Split(startString, "-")
		// split string [] string
		subStringLeft := strings.Split(string(newString[0]), ".")
		// split string [] string
		subStringRigth := strings.Split(string(newString[1]), ".")
		// CRUTCH , take subStringLeft[2]  static value
		twoIndexStaticNumber, _ := strconv.Atoi(subStringLeft[2])

		fmt.Sprintln(subStringRigth) // compiler scold, not use variable
		fmt.Sprintln(subStringLeft)  // compiler scold, not use variable

		//  convert  subStringLeft and subStringRigth []int, equals
		//  index value and  append  in []int differenceBetweenIndex
		for i := 2; i < len(subStringLeft); i++ {

			if subStringLeft[i] != subStringRigth[i] || subStringLeft[i] == subStringRigth[i] {

				left, _ := strconv.Atoi(subStringLeft[i])
				right, _ := strconv.Atoi(subStringRigth[i])

				differenceBetweenIndex = append(differenceBetweenIndex, -(left - right))
				fmt.Println(differenceBetweenIndex)
			}

		}

		for _, val := range subStringLeft {

			value, _ := strconv.Atoi(val)

			convertStringSliceToInt = append(convertStringSliceToInt, value)
		}

		for i = convertStringSliceToInt[2]; i <= (convertStringSliceToInt[2] + differenceBetweenIndex[0]); i++ {
			fmt.Println(convertStringSliceToInt)
			fmt.Println(differenceBetweenIndex)
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
	defer file.Close()
}
