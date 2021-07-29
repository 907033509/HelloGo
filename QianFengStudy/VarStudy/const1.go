package main

import "fmt"

func main() {
	/**
	常量：
	1.概念：同变量类似，程序执行过程中数值不能改变
	2.语法：
		显式类型定义
		隐式类型定义
	3.常数：
		固定数值：100，"abc"
	4.变量必须使用 常量不一定需要使用
	5.常量定义以后，不允许在修改数值
	6.变量中的数据类型只可以是布尔型，数字型（整数型，浮点型，复数）和字符串型
	7.显示指定类型的时候，必须确保常量左右值一致，需要时可做显示类型转换。变量是可以不同的类型值
	*/
	fmt.Println(100)
	fmt.Println("hello")

	//1.定义常量
	const PATH string = "http:www.baidu.com"
	const PI = 3.14
	fmt.Println(PATH)

	//2.尝试修改变量的数值：
	// PATH = "hello" //cannot assign to PATH

	//3.定义一组常量
	const C1, C2, C3 = 100, 3.14, "HHH"
	const (
		A1 = 1
		A2 = 2
		A3 = 3
	)

	//4.一族常量中，如果某个常量没有初始值，默认和上一行一致
	const (
		a int = 100
		b
	)
	fmt.Println(a)
	fmt.Println(b)

	//5.枚举类型：使用常量作为枚举类型。一组相关数值的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
