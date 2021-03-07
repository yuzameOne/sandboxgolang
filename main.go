package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"sync"
)

var procsThreads int = runtime.GOMAXPROCS(0)
var rangeIPStartFile []string
var chanRangeIPStartFile = make(chan string)
var countRangeRangeSliceIndex int
var mu sync.Mutex

func main() {

	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("file not create : %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		rangeIPStartFile = append(rangeIPStartFile, scanner.Text())
	}

	for i := 0; i < len(rangeIPStartFile); i++ {

		go func() {

			mu.Lock()
			countRangeRangeSliceIndex = countRangeRangeSliceIndex + 1
			mu.Unlock()
			
			
		}()
	}

}
