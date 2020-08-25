package main

import (
	"fmt"
)

type celsius float64
type kelvin float64
type fahrenheit float64

func main() {
	var c celsius = 36.6
	var k kelvin = 309.75
	var f fahrenheit = 97.88
	var cf fahrenheit = c.celsiusToFahrenheit()
	var ck kelvin = c.celsiusTokelvin()
	var kc celsius = k.kelvinToCelsius()
	var kf fahrenheit = k.kelvinToFahrenheit()
	var fc celsius = f.fahrenheitToCelsius()
	var fk kelvin = f.fahrenheitTokelvin()
	fmt.Printf("摄氏度(%.2f)-->华摄氏度(%.2f)\n", c, cf)
	fmt.Printf("摄氏度(%.2f)-->开尔文(%.2f)\n", c, ck)
	fmt.Printf("开尔文(%.2f)-->摄氏度(%.2f)\n", k, kc)
	fmt.Printf("开尔文(%.2f)-->华摄氏度(%.2f)\n", k, kf)
	fmt.Printf("华摄氏度(%.2f)-->摄氏度(%.2f)\n", f, fc)
	fmt.Printf("华摄氏度(%.2f)-->开尔文(%.2f)\n", f, fk)
}

//kelvinToCelsius celsiusToFahrenheit kelvinToFahrenheit
// 摄氏度-->华摄氏度
func (c celsius) celsiusToFahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// 摄氏度-->开尔文
func (c celsius) celsiusTokelvin() kelvin {
	return kelvin(c + 273.15)
}

// 开尔文-->摄氏度
func (k kelvin) kelvinToCelsius() celsius {
	return celsius(k - 273.15)
}

// 开尔文-->华摄氏度
func (k kelvin) kelvinToFahrenheit() fahrenheit {
	return fahrenheit(k.kelvinToCelsius().celsiusToFahrenheit())
}

// 华摄氏度-->摄氏度
func (f fahrenheit) fahrenheitToCelsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

// 华摄氏度-->开尔文
func (f fahrenheit) fahrenheitTokelvin() kelvin {
	return kelvin(f.fahrenheitToCelsius().celsiusTokelvin())
}
