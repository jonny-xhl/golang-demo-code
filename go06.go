package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("浮点数学习")
	// 单精度 4字节
	var pi1 float32 = math.Pi
	// 双精度 8字节
	var pi2 float64 = math.Pi
	fmt.Printf("float32:%v float64:%v\n", pi1, pi2)

	// 零值
	var zero64 float64
	var zero32 float32
	zero := 0.0
	fmt.Printf("zero64:%v zero32:%v zero:%v\n", zero64, zero32, zero)

	// 格式化打印
	// Print Println打印浮点数时尽可能多的显示几位小数
	fmt.Print(pi2)
	fmt.Printf("\n%v\n", pi2)
	fmt.Printf("%f\n", pi2)
	fmt.Println(pi2)
	fmt.Println("===%f格式化输出===")
	p := 1.0 / 3
	fmt.Print(p)
	fmt.Println()
	fmt.Println(p)
	fmt.Printf("%f\n", p)
	fmt.Printf("%.3f\n", p)   // 保留三位小数
	fmt.Printf("%5.2f\n", p)  // 总共4位宽度，2位小数，左补空格
	fmt.Printf("%05.2f\n", p) // 总共4位宽度，2位小数，左补零。经过测试宽度包含了小数点（.）

	// 浮点类型的精度
	fmt.Println("===浮点类型的精度===")
	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)
	// 总结：浮点数类型不适合用于金融类计算，为了尽量最小化舍人错误，建议先做乘法，再做除法
	// 测试
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32.0, "℉\n")
	fmt.Print((celsius/5.0*9.0)+32.0, "℉\n")
	fmt.Print((celsius*9.0/5.0)+32.0, "℉\n")

	// 如何比较浮点数类型，精度比较如何避免
	fmt.Println(piggyBank == 0.3)
	// 使用万分之精度来进行比较，这个和js\c\c#语言中都存在这个问题
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)

	fmt.Println("======作业题======")
	homework06()
}

const totle06 = 20

func homework06() {
	fmt.Println("======go06作业题======")
	var coin05, coin01, coin025 float64 = 0.05, 0.10, 0.25
	// r := rand.Intn(2) + 1
	var deposit float64
	for deposit < totle06 {
		switch rand.Intn(3) + 1 {
		case 1:
			deposit += coin01
		case 2:
			deposit += coin025
		case 3:
			deposit += coin05
		}
		fmt.Printf("存钱罐的余额为：%.2f\n", deposit)
	}
}
