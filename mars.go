package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 单个变量
	var num = 123
	fmt.Println(num)
	// 多个变量 第一种
	var num1, num2 = 123, 456
	fmt.Println(num1, num2)

	// 第二种
	var (
		num3 = 123
		num4 = 456
	)
	fmt.Println(num3, num4)

	var num5 = 123
	var num6 = 456
	fmt.Println(num5, num6)

	// 输出格式化
	// 6位长度，左边补空格
	fmt.Printf("%6v\n", num1)
	// 负数右补齐
	fmt.Printf("%-6v %3v\n", num1, num2)

	// const
	const num7 = 123
	fmt.Println(num7)

	// compute
	const num8 = 123
	const num9 = num8 * 11.0
	const num10 = num9 / 1.7
	fmt.Println(num10)

	// rand
	var num11 = rand.Intn(10) + 1
	fmt.Println(num11)
	num11 = rand.Intn(10) + 1
	fmt.Println(num11)

	// 作业
	// 编写程序来确定飞船要在28天内到达Malacandra的行进速度（公里/小时），假设距离为56,000,000公里。
	const distance, day = 56000000, 28
	fmt.Printf("飞船的速度为：%vkm/h\n", getSpeed1(distance, day))
}

func getSpeed(distance, day int) int {
	// 天转小时
	var hour = day * 24
	return distance / hour
}

func getSpeed1(distance, day int) (speed int) {
	// 天转小时
	var hour = day * 24
	speed = distance / hour
	return speed
}
