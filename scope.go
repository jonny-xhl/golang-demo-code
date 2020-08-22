package main

import (
	"fmt"
	"math/rand"
)

// package级别作用域
// var PValue = 123

// 作用域  短声明练习
func main() {
	for num := 1; num <= 10; num++ {
		fmt.Printf("产生随机日期-%v：", num)
		getDate()
	}
}

// 实现随机生成年月日
func getDate() {
	year := rand.Intn(50) + 2020
	// 随机哪月
	mounth := rand.Intn(12) + 1
	// 默认每月31天
	day := 31
	switch mounth {
	case 2:
		if isLeapYear(year) {
			day = 29
		} else {
			day = 28
		}
	case 4, 6, 9, 11:
		day = 30
	}
	day = rand.Intn(day) + 1
	printRandDate(year, mounth, day)
}

func isLeapYear(year int) (isLeap bool) {
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		isLeap = true
	} else {
		isLeap = false
	}
	return isLeap
}

func printRandDate(year, mounth, day int) {
	fmt.Printf("%04v/%02v/%02v\n", year, mounth, day)
}
