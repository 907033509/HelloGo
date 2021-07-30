package main

import (
	"fmt"
)

func main(){
/*
字符串：
1.概念：多个byte的集合，理解为一个字符序列
2.语法：使用双引号（也可以使用单引号）
	"abc","hello"
3.转义字符： \
	a:有一些字符有特殊的作用，可以转义为不同字符
	b:有一些字符，就是一个普通字符，转义后有特殊的作用
*/

//1.定义字符串
var s1 string
s1 = "王二狗"
fmt.Printf("%T,%s\n",s1,s1)

s2 := 'A'
s3 := "A"
fmt.Printf("%T,%s\n",s2,s2)
fmt.Printf("%T,%s\n",s3,s3)

s4 := '中'
fmt.Printf("%T,%d,%c,%q\n",s4,s4,s4,s4)

//2.转义字符
fmt.Println("\"hello world\"")
fmt.Println("hello wor\tld")
}