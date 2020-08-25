package main

import (
	"fmt"
)

func main() {
	k := 294.0
	celsius := kelvinToCelsius(k)
	fmt.Println(celsius, "℃")
	//作业题
	k0 := 0.0
	fmt.Println(kelvinToFahrenheit(k0), "℉")
}

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	// K-->℃-->℉
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

// 任意多个参数
func params(a ...interface{}) {

}

// 多个入参
func manyParameter(a, b int, s string) {

}

// 多个出参
func manyRerutn() (a, b int, s string) {
	return a, b, s
}
