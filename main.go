package main

import (
	"fmt"
	"strings"
)

var newCustomString strings.Builder
var subSlice []string

func main() {

	one, two := "0", "1"

	fmt.Printf("one : %s, two : %s  \n", one, two)

	fmt.Println(&one, &two)

	one, two = "2", "3"

	fmt.Printf("one : %s, two : %s  \n", one, two)

	fmt.Println(&one, &two)

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer f.Close()

	// r := bufio.NewReader(f)

	// for {
	// 	line, err := r.ReadString('\n')

	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	stringRange = append(stringRange, line)

	// }

	// fmt.Println(stringRange)

	// str := "5.100.67.0-5.100.67.255"

	// for index, value := range str {

	// 	if value == 46 || value == 45 {

	// 		subSlice = append(subSlice, newCustomString.String())
	// 		newCustomString.Reset()
	// 	}

	// 	if value != 46 && value != 45 {

	// 		newCustomString.WriteByte(str[index])

	// 		if len(str)-1 == index {
	// 			subSlice = append(subSlice, newCustomString.String())
	// 		}

	// 	}

	// }
	// fmt.Println(subSlice[1])

	// index 2,3 and 6,7
	// нужно в for сравнивать СТРОКИ index 2,3 and 6,7 если меньше  конвертировать в int, инкрементить
	// и тут же конверитить в стороку и собирать это в нову строку StringBulder

	//  СУКА 2 фукции 3 канала)))))
}

// func removesPointAndDash(arrayRange []string) {

// 	str := "5.100.67.0-5.100.67.255"

// 	point := bytes.Index([]byte(str), []byte("-"))

// 	leftString := str[:point]
// 	rightString := str[point+1:]

// 	fmt.Println(leftString)
// 	fmt.Println(rightString)

// 	// TODO
// 	//  что надо  67.0 и  67.255 из стиринги в  инты сравить и обратно в строку  bytes.Buffer

// 	// 21.10 14:11 стоял на смене и пришла мысль в голову , что последнее число  всегда будет в диапапозоне
// 	//  от 0 до 255 (67.0). нужно просто создать Enum от 0 до 255 и подставить в цикле в строку  bytes.Buffer
// }

// 22.10 3:06  81 это (56 49) а 82 это (56 50). 112 (49 49 50) - 123(49 50 51).
//  циферки в таблице ascii от 0 до 9  в таблице ascii  от 48  до 57
//  ТО ЕСТЬ  увеличиваем крайний байт до  0 (112 - 119  потом  сбрасываем на 0(48) 110) с увеличением предыдущего
// байта на +1 (110 [49 49 48] до 120 [49 50 48])
