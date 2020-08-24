package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//类型转换
	fmt.Println("========类型转换========")

	// integer to string
	num := 97
	str := string(num)
	fmt.Println("convert integer to string use string() and value is:", str)
	str1 := strconv.Itoa(num)
	fmt.Println("convert integer to string use strconv.Itoa() value is:", str1)
	str2 := fmt.Sprintf("%v", num)
	fmt.Println("convert integer to string use fmt.Sprintf() value is:", str2)

	// string to integer
	str3 := "56" // 例如这里 56S就会转换错误
	num1, err := strconv.Atoi(str3)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("convert string to integer use strconv.Atoi() and value is:", num1)

	// int to float
	num3 := 100
	f1 := float32(num3)
	f2 := float64(num3)
	fmt.Println(f1, " ", f2)

	// float to int
	f3 := 129.123
	num4 := int(f3)
	fmt.Println(num4)
	num5 := int8(f3)
	fmt.Println("int8:", num5) // 这里超出了范围转换时会出现“环绕”情况，在类型转换的时候需要注意这种情况，需要逻辑判断
	num6 := int16(f3)
	fmt.Println("int16:", num6)
	num7 := int64(f3)
	fmt.Println("int64:", num7)
	num8 := uint8(f3)
	fmt.Println("uint8:", num8)
	num9 := uint16(f3)
	fmt.Println("uint16:", num9)
	num10 := uint(f3)
	fmt.Println("uint:", num10)
	num11 := uint64(f3)
	fmt.Println("uint64:", num11)

	// overflow handle
	if f3 <= math.MaxInt8 && f3 >= math.MinInt8 {
		num12 := int8(f3)
		fmt.Println("overflow:", num12)
	} else {
		num12 := int16(f3)
		fmt.Println("overflow:", num12)
	}

	// int to other int
	var num13 int = 100
	num14 := uint8(num13)
	fmt.Println("int to uint:", num14)

	// bool to string
	b := true
	s := fmt.Sprintf("%v", b)
	fmt.Println("bool to string:", s)
	var s1 string
	if b {
		s1 = "true"
	} else {
		s1 = "false"
	}
	fmt.Println("bool to string:", s1)

	//作业题
	homework10()
}

func homework10() {
	fmt.Println("========go10作业题========")
	if convertStrToBoolean("true") || convertStrToBoolean("yes") || convertStrToBoolean("1") {
		fmt.Printf("========%v========\n", true)
	}
}

func convertStrToBoolean(str string) bool {
	if str == "true" || str == "yes" || str == "1" {
		return true
	} else if str == "false" || str == "no" || str == "0" {
		return false
	} else {
		fmt.Println("转换错误")
		return false
	}
}
