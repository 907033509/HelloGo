package main

import "fmt"

var a = 1000 //全局变量
var b = 2000
// c := 3000 语法错误
func main(){
	/*
	注意点：
	1.变量必须定义才可以使用
	2.变量的类型和赋值必须一致
	3.同一个作用域，变量名不能冲突
	4.简短定义的方式，左边的变量至少有一个是新的
	5.简短定义方式，不能定义全局变量
	6.变量的零值（也就是默认值）
		整型：默认值是 0
		浮点型：默认是 0（本质是0.0）
		字符串：默认是 " "
	*/
	var num int
	num = 100
	fmt.Printf(" Printf 的输出："  + "num的数值是：%d，地址是：%p\n",num,&num)

	num = 200
	fmt.Printf(" Printf 的输出："  + "num的数值是：%d，地址是：%p\n",num,&num)

	fmt.Println("=======================================")
	var a int // 整数 默认值是0
	fmt.Println(a)
	var s float64 // 0.0 
	fmt.Println(s)
	var d string // ""
	fmt.Println(d)
	var f []int 
	fmt.Println(f)
	fmt.Println(f==nil) 

}
