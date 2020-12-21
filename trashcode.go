// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"strings"
// )

// var newString strings.Builder
// var count int
// var some string
// var newLine byte = '\n'
// var hyphen byte = '-'
// var copyString string
// var lastNumber int

// func main() {

// 	file, err := ioutil.ReadFile("unparseip.txt")

// 	if err != nil {
// 		log.Printf("file dont read : %s", err)
// 	}

// 	for _, v := range file {

// 		newString.WriteByte(v)

// 		if v == newLine {

// 			s := fmt.Sprintf("%s", &newString)

// 			// fund last number  in string

// 			for i := 0; i < len(s); i++ {

// 				endOfString := s[0 : len(s)-1]

// 				fmt.Println(endOfString)
// 				break

// 			}

// 			// for i := 0; i < len(s); i++ {
// 			// 	fmt.Printf("%v \n", s[i])
// 			// 	fmt.Printf("string data : %s \n", string(s[i]))
// 			// 	if s[i] == 10 {
// 			// 		lastNumber = s[1 : len(s)-1]
// 			// 		fmt.Printf(" print last  number string in loop : %s\n", lastNumber)
// 			// 	} else if s[i] == 45 {
// 			// 		copyString = s[0:i]
// 			// 		fmt.Printf(" print string in loop : %s\n", copyString)
// 			// 	}
// 			// }

// 			// str := strings.NewReader(copyString)

// 			// fmt.Printf("length string  copyString : %d \n", str.Len())
// 			// fmt.Printf("last index  copyString : %s", string(copyString[9]))

// 			newString.Reset()

// 			count = count + 1

// 		} else if count == 5 {
// 			break
// 		}

// 	}

// 	fmt.Println(some)

// }
