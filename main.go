package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

var chanInt = make(chan int)
var rangeIp = make(map[int]string)
var counMapIndex int
var mu sync.Mutex

// TODO
// написать benchmark на скорость
// добавить в файл до 100 строк example.txt
// fan in / fan out  посмотреть отличие
// написать benchmark на скорость и сравнить на скорость

func main() {

	//  open file
	file, err := os.Open("example.txt")

	if err != nil {
		log.Fatal("file not open", err)
	}

	//  ...
	scanner := bufio.NewScanner(file)
	//  ...
	scanner.Split(bufio.ScanLines)

	//  scan and write in map  string
	for scanner.Scan() {
		//  iterate  int  index map
		counMapIndex = counMapIndex + 1

		//  write in map  string ip range
		rangeIp[counMapIndex] = scanner.Text()
	}

	defer file.Close()

	// TODO
	// адаптировать этот кусок под этот код( хочу спать не могу сформулировать правильно мысль )
	// засунуть в горутины (анонимную)

	// for i := 0; i < len(parseIPLinesFile); i++ {

	// 	// take string like  "5.100.67.0-5.100.67.255"
	// 	startString := parseIPLinesFile[i]
	// 	//  split string in []string
	// 	newString := strings.Split(startString, "-")
	// 	// split string [] string
	// 	subStringLeft := strings.Split(string(newString[0]), ".")
	// 	// split string [] string
	// 	subStringRigth := strings.Split(string(newString[1]), ".")
	// 	// CRUTCH , take subStringLeft[2]  static value
	// 	twoIndexStaticNumber, _ := strconv.Atoi(subStringLeft[2])

	// 	fmt.Sprintln(subStringRigth) // compiler scold, not use variable
	// 	fmt.Sprintln(subStringLeft)  // compiler scold, not use variable

	// 	//  convert  subStringLeft and subStringRigth []int, equals
	// 	//  index value and  append  in []int differenceBetweenIndex
	// 	for i := 2; i < len(subStringLeft); i++ {

	// 		if subStringLeft[i] != subStringRigth[i] || subStringLeft[i] == subStringRigth[i] {

	// 			left, _ := strconv.Atoi(subStringLeft[i])
	// 			right, _ := strconv.Atoi(subStringRigth[i])

	// 			differenceBetweenIndex = append(differenceBetweenIndex, -(left - right))

	// 		}

	// 	}

	// 	//  delete all elements slice convertStringSliceToInt
	// 	if len(convertStringSliceToInt) != 0 {
	// 		for i := 0; i < 4; i++ {

	// 			convertStringSliceToInt = convertStringSliceToInt[:len(convertStringSliceToInt)-1]
	// 		}
	// 	}

	// 	//  convert all string elements to int type , and append  to slice convertStringSliceToInt
	// 	for _, val := range subStringLeft {

	// 		value, _ := strconv.Atoi(val)

	// 		convertStringSliceToInt = append(convertStringSliceToInt, value)

	// 	}

	// 	// open/create file where write finaly range ip
	// 	file, err := os.OpenFile(argTwo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	// 	if err != nil {
	// 		log.Fatalf("file not create : %s", err)
	// 	}
	// 	// get three index to int slice  and increment him
	// 	for i := convertStringSliceToInt[2]; i <= (convertStringSliceToInt[2] + differenceBetweenIndex[0]); i++ {

	// 		if i == (twoIndexStaticNumber+differenceBetweenIndex[0])+1 {
	// 			break
	// 		}
	// 		convertStringSliceToInt[2] = i

	// 		// get get index to int slice  and increment him
	// 		for j := convertStringSliceToInt[3]; j <= (convertStringSliceToInt[3] + differenceBetweenIndex[1]); j++ {

	// 			convertStringSliceToInt[3] = j

	// 			//  convert all int elements to string type , and append  to slice convertStringSliceToInt
	// 			for _, val := range convertStringSliceToInt {

	// 				value := strconv.Itoa(val)

	// 				convertIntSliceToString = append(convertIntSliceToString, value)
	// 			}
	// 			// convert string slice to string
	// 			finishStringWriteToFile = strings.Join(convertIntSliceToString, ".")

	// 			// delete all element slice sitring convertIntSliceToString
	// 			for i := 0; i < len(convertIntSliceToString); {

	// 				convertIntSliceToString[len(convertIntSliceToString)-1] = " "
	// 				convertIntSliceToString = convertIntSliceToString[:len(convertIntSliceToString)-1]
	// 			}
	// 			// write finish string to file
	// 			file.WriteString(finishStringWriteToFile + "\n")

	// 			// counter string
	// 			count = count + 1

	// 			// console output number string line in the file
	// 			fmt.Printf("lines in file : %d \n", count)

	// 			//  if j == 255
	// 			if j == differenceBetweenIndex[1] {
	// 				convertStringSliceToInt[3] = 0

	// 				break
	// 			}

	// 		}
	// 		file.Close()
	// 	}

	// for i := 0; i < 100; i++ {
	// 	go func() {

	// 		mu.Lock()
	// 		x = x + 1
	// 		chanInt <- x
	// 		mu.Unlock()
	// 	}()

	// 	fmt.Println(<-chanInt)
	// }

	// close(chanInt)
}
