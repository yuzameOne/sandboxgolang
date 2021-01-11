package read

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Read() {

	file, err := ioutil.ReadFile("unparseip.txt")

	if err != nil {

		log.Fatal("file not open", err)
	}
	fmt.Println(string(file))
}
