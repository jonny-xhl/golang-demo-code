package main

import (
	"fmt"
	"math/rand"
)

// 《Get Programming with Go》 || 《Go 语言趣学指南》第一部分习题

const distance = 62100000

func main() {
	fmt.Printf("%-20v %-10v %-20v %-10v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("===================================================================")
	for num := 1; num <= 10; num++ {
		createMileage()
	}
}

// 输出一条里程
func createMileage() {
	speed := getSpeed()
	days := distance / (speed * 24)
	tripTypeNum := getTripType()
	tripType := "Unknown"
	switch tripTypeNum {
	case 1:
		tripType = "One-way"
	case 2:
		tripType = "Round-trip"
	}
	price := tripTypeNum * getPrice()
	spaceline := getSpaceline()
	fmt.Printf("%-20v %-10v %-20v $ %-10v\n", spaceline, days, tripType, price)
}

func getSpaceline() (spaceline string) {
	switch line := rand.Intn(3) + 1; line {
	case 1:
		spaceline = "Space Adventures"
	case 2:
		spaceline = "SpaceX"
	case 3:
		spaceline = "Virgin Galactic"
	}
	return spaceline
}

// 获取速度；返回km/h
func getSpeed() int {
	return (rand.Intn(30-16) + 1) * 60 * 60
}

//  获取里程类型；1-单程  2-往返
func getTripType() int {
	return rand.Intn(2) + 1
}

// 票价；单位：百万美元（单程）
func getPrice() int {
	return rand.Intn(5000-3600) + 1
}
