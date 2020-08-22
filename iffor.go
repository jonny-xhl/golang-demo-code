package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("========if练习========")
	// if
	const name = "jonny xiang"
	if name == "jonny xiang" {
		fmt.Println("if:my name is jonny xiang")
	} else {
		fmt.Println("if:my name is not jonny xiang")
	}

	// switch
	fmt.Println("========switch无值练习========")
	switch {
	case name == "jonny xiang":
		fmt.Println("switch:my name is jonny xiang")
		fallthrough // 这里存在疑惑
	case !strings.Contains(name, "xiang"):
		fmt.Println("switch:my name contains xiang")
		fallthrough
	case "a" == "":
		fmt.Println("fallthrough尝试")
	}
	fmt.Println("========switch有值练习========")
	switch name {
	case "jonny xiang":
		fmt.Println("switch:my name is jonny xiang")
		fallthrough // 这里存在疑惑
	default:
		fmt.Println()
	}

	// for
	fmt.Println("========开始for练习========")
	var count = 10
	for count > 0 {
		fmt.Println(count)
		// time.Second为1s, ns:n*time.Second
		// time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 500)
		count--
	}

	count1 := 10
	fmt.Printf(":=赋值练习测试,count1:%v", count1)
}
