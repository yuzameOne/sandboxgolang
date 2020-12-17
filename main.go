package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var newString strings.Builder
var count int

func main() {

	file, err := ioutil.ReadFile("unparseip.txt")

	if err != nil {
		log.Printf("file dont read : %s", err)
	}

	for _, v := range file {

		newString.WriteByte(v)

		// 10 == "\n"
		if v == 10 {

			some := fmt.Sprintln(newString.String())

			fmt.Print(string(some[9]))
			fmt.Print(string(some[20]))
			fmt.Print(string(some[21]))
			fmt.Print(string(some[22]))
			// fmt.Print(string(some[23]))
			//  23 == "\n"

			// for key, val := range some {

			// 	// 45 == "-"
			// 	if val == 45 {
			// 		key = key - 1
			// 		fmt.Printf("key index : %v\n ", key)
			// 	} else if val == 10 {
			// 		key = key - 1
			// 		fmt.Printf("key index : %v\n", key)
			// 		break
			// 	}
			// }
			newString.Reset()

			count = count + 1

		} else if count == 1 {
			break
		}

	}

}
