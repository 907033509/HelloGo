package main

import "fmt"

func main(){
	/*
	变量：variable
	概念：一小块内存，用于存储数据结构，在程序运行过程中数值可以改变
	使用：
		step1: 变量的声明，也叫作定义
			第一种：var 变量名字 数据类型
				var 变量名 数据类型 = 赋值
			第二种：类型推断，省略数据类型
				var 变量名 = 赋值
			第三种：简短声明，省略var
				变量名 := 赋值
		step2：变量的访问，也叫作赋值和取值
			直接根据变量名访问
	*/
	
	//第一种：定义变量，然后赋值
	var num1 int = 30
	fmt.Println("num1的数值是：%d\n",num1)

	//第二种：定义变量，然后赋值
	var name1 = "王二狗"
	fmt.Println("name1的数值是：%d\n",name1)

	//第三种：剪短定义，也叫作简短声明
	sum1 := 100
	fmt.Println("sum1的数值是：%d\n",sum1)

	//多个变量同时定义
	var a,b,c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a,b,c)

	var m,n int = 100,200
	fmt.Println(m,n)

	var b1,n1,m1 = 1,2,"lihua"
	fmt.Println(b1,n1,m1)

	var (
		name = "李晓华"
		age = 18
		sex = "女"
	)
	fmt.Println(name,age,sex)
}
