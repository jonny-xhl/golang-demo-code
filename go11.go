package main

import (
	"fmt"
	"strings"
)

const secret = "GOLANG"

func main() {
	input := "ihioqwhiwqeiohiowqe"
	// fmt.Printf("%c\n", 90)
	fmt.Println(decodeInput(input))
}

func decodeInput(str string) (result string) {
	str = strings.ToUpper(str)
	var ch int
	var res rune
	for i, c := range str {
		//密钥第几位
		ch = i % 6
		res = c + rune(secret[ch]-65)
		// fmt.Println("%c", res)
		// if res > 'Z' {
		// 	res = res - 90
		// }
		result += string(res)
	}
	return result
}
