package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("姓名：向洪林\n年龄：25")
	fmt.Println(`姓名：向洪林\n年龄：25`) // 使用``进行转义输出
	fmt.Println(`姓名：向洪林
年龄：25`)

	goDir := `C:\go`
	fmt.Println("go安装路径:", goDir)
	fmt.Println("go安装路径:C:\\go")

	// unicode
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	// 输出code point
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	// 自定义类型
	type cfloat = float32
	type crune = int
	var cf cfloat = 1.0
	fmt.Printf("%.2f\n", cf)

	str := "my name is jonny 向"
	//len
	fmt.Println("str bytes len:", len(str))
	fmt.Println("str len:", utf8.RuneCountInString(str))
	r, size := utf8.DecodeRuneInString(str)
	fmt.Printf("char:%c code:%v size:%v", r, r, size)

	//range
	for _, c := range str {
		fmt.Printf("%c\n", c)
	}

	// 作业题
	homework09Part1()
	fmt.Println("========作业题2=======")
	homework09Part2()
}

func homework09Part1() {
	const str = "L fdph,L vdz,L frqtxhuhg"
	var decodeStr string
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]-3)
	}
	for i := 0; i < utf8.RuneCountInString(str); i++ {
		decodeStr += string(str[i] - 3)
	}
	fmt.Println("decode str value is ", decodeStr)
}

func homework09Part2() {
	const str = "Hola Estación Espacial Internacional"
	var encodeStr string
	for _, c := range str {
		if c < 'A' || c > 'z' {
			encodeStr += string(c)
		} else {
			encodeStr += string(c + 13)
		}
	}
	fmt.Println("encode str value is ", encodeStr)
}
