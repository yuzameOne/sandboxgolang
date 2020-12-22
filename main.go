package main

import "fmt"

var stringLength int

func main() {
	s := "80.92.160.0-80.92.175.255"

	// find full  string length

	for i := 0; i < len(s); i++ {

		stringLength = i
		// fmt.Println(stringLength)

	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v \n", s[i])
		fmt.Printf("string data : %s \n", string(s[i]))
		if s[i] == 10 {
			lastNumber = s[1 : len(s)-1]
			fmt.Printf(" print last  number string in loop : %s\n", lastNumber)
		} else if s[i] == 45 {
			copyString = s[0:i]
			fmt.Printf(" print string in loop : %s\n", copyString)
		}
	}

}
