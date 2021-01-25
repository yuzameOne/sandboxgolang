package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	var argOne string
	flag.StringVar(&argOne, "pathtofile", "", "path to file ")

	if argOne != "" {
		fmt.Println("the first argument is missing")
		os.Exit(3)
	}

	var argTwo string
	flag.StringVar(&argTwo, "pathtosavefile", "", "path to file ")

	flag.Parse()

	if argTwo == "" {
		argTwo = path + "/new_ip_" + argOne
	}

	fmt.Println(argOne)
	fmt.Println(argTwo)
}
