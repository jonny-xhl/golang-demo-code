package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("========较大的数练习========")
	var bigNum = 10e28
	fmt.Println(bigNum)
	//var bigNum1 uint64 = 10e28
	//fmt.Println(bigNum1)

	a := big.NewInt(465999)
	b := big.NewFloat(646565.236)
	c := big.NewRat(1, 3)
	fmt.Println(a, b, c)

	var distance = new(big.Int)
	distance.SetString("2400000000000000000000", 10)
	fmt.Println(distance)

	d := big.NewInt(123456789)
	e := big.NewInt(100)
	f := new(big.Int)
	fmt.Println("int.Div", f.Div(d, e))
	// 作业题
	homework08()
}

func homework08() {
	fmt.Println("========go09作业题========")
	const distance = 236000000000000000 * 1000
	fmt.Printf("%.2f光年\n", distance/lightYears())
}

func lightYears() float64 {
	//光速299792458m/s 平均一年天数365.25
	const lightSpeed = 299792458
	seconds := 365.25 * 24 * 60 * 60 * 60
	return lightSpeed * seconds
}
