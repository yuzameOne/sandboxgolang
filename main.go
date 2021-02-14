package main

import (
	"bufio"
	"log"
	"os"
)

var rangeIp = make(map[int]string)
var counMapIndex int

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

		counMapIndex = counMapIndex + 1

		rangeIp[counMapIndex] = scanner.Text()
	}

	defer file.Close()

	// for key, val := range rangeIp {
	// 	fmt.Printf("key %d , value %s \n", key, val)
	// }
}
