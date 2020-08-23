package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 整数练习
func main() {
	// 所有“整数”类型
	var nuint uint
	var nint int
	var nint8 int8
	var nint16 int16
	var nint32 int32
	var nint64 int64
	var nuint8 uint8
	var nuint16 uint16
	var nuint32 uint32
	var nuint64 uint64
	fmt.Printf("%T\n", nuint)
	fmt.Printf("%T\n", nint)
	fmt.Printf("%T\n", nint8)
	fmt.Printf("%T\n", nint16)
	fmt.Printf("%T\n", nint32)
	fmt.Printf("%T\n", nint64)
	fmt.Printf("%T\n", nuint8)
	fmt.Printf("%T\n", nuint16)
	fmt.Printf("%T\n", nuint32)
	fmt.Printf("%T\n", nuint64)

	// %T %x %b
	num := 10
	fmt.Printf("type:%T | hexadecimal:%x binary:%b\n", num, num, num)
	// 补0
	fmt.Printf("type:%T | hexadecimal:%08x binary:%08b\n", num, num, num)

	// 溢出情况
	// 第一种：无符号
	var unum8 uint8 = 255
	fmt.Printf("unum8:%v | binary:%08b\n", unum8, unum8)
	unum8++
	fmt.Printf("unum8:%v | binary:%08b\n", unum8, unum8)
	// 第二种：有符号
	var num8 int8 = 127
	fmt.Printf("num8:%v | binary:%08b\n", num8, num8)
	num8++
	fmt.Printf("num8:%v | binary:%08b\n", num8, num8)

	// %T输出类型
	a := 123
	b := 1.23
	c := true
	d := "text"
	fmt.Printf("Type %T for %[1]v\n", a)
	fmt.Printf("Type %T for %[1]v\n", b)
	fmt.Printf("Type %T for %[1]v\n", c)
	fmt.Printf("Type %T for %[1]v\n", d)

	// unix时间 1970-1-1（UTC/GMT的午夜）开始所经过的秒数
	fmt.Print(time.Now(), "\n")
	future := time.Unix(1598184411, 0)
	fmt.Println(future)

	// math自带的整数极限值
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxUint32)
	// 这里会异常,超出int范围
	//fmt.Println(math.MaxUint64)

	// 作业题
	homework07()
}

const totle07 uint16 = 20 * 100

func homework07() {
	fmt.Println("========go07作业题========")
	var coin5, coin10, coin25 uint16 = 5, 10, 25
	var deposit uint16
	for deposit < totle07 {
		switch rand.Intn(3) + 1 {
		case 1:
			deposit += coin25
		case 2:
			deposit += coin10
		case 3:
			deposit += coin5
		}
		fmt.Printf("存钱罐的余额为：$%.2f\n", float32(deposit)/100.0) // 这里引入了类型转换
	}
}
