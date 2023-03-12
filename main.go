package main

import (
	"bytes"
	"fmt"
	"strconv"
)

var buffer bytes.Buffer
var sm = make(map[string]int)

func transformString(s string) {

	for _, v := range s {

		value, exist := sm[string(v)]

		if exist {
			sm[string(v)] = value + 1
		} else {
			sm[string(v)] = 1
		}
	}

}

func main() {

	transformString("asdasdasdbffagdkjkirktgnjnttt")

	for k, v := range sm {
		buffer.WriteString(k)
		s := strconv.Itoa(v)
		buffer.WriteString(s)
	}

	fmt.Println(buffer.String())

}
