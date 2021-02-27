package main

import (
	"fmt"
)

func main() {
	str := []rune("apple你好")
	str2 := string(reverse(str))
	fmt.Printf("before: %+v\n", string(str))
	fmt.Printf("after : %+v\n", str2)

	str3 := "apple你好"
	fmt.Printf("after : %+v\n", reversev2(str3))

	str4 := "apple你好"
	fmt.Printf("after : %+v\n", reversev3(str4))

	str5 := "apple"
	fmt.Printf("after : %+v\n", reversev3(str5))

}

// reverse a string
func reverse(str []rune) []rune {
	var returnArr []rune
	for i := len(str) - 1; i >= 0; i-- {
		returnArr = append(returnArr, str[i])
	}
	return returnArr
}

// reverse a string
func reversev2(str string) string {
	var str2 []rune
	fmt.Printf("%+v %d \n", str, len(str))
	for _, i := range str {
		fmt.Printf("%+v\n", i)
		str2 = append(str2, rune(i))
	}
	fmt.Printf("%+v %d \n", str2, len(str2))
	var returnArr []rune
	for i := len(str2) - 1; i >= 0; i-- {
		returnArr = append(returnArr, str2[i])
	}
	return string(returnArr)
}

// reverse a string
func reversev3(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
