package main

import (
	"fmt"
)

func main(){
	/*
	1.基本数据类型：
		布尔类型：bool
			取值：true,false
		数值类型：
			整数：int
				有符号：最高位表示符号，0整数，1复数，其余为表示数值
					int8
					int16
					int32
					int64
				无符号：
					uint8
					uint16
					uint32
					uint64

				int,uint
				byte:uint8
				rune:int32
			浮点：生活中的小数
			复数：complex
		字符串：string
	2.复合数据类型：
		array,slice,amp,function,pointer...
	*/

	//1.布尔类型
	var b1 bool
	b1 = true
	fmt.Println("%T,%t\n",b1,b1)
	var a uint8
	var b byte
	fmt.Println(a,b)

	//浮点
	var f1 float32
	f1 = 3.14
	var f2 float64
	f2 = 3.14
	fmt.Println("%T,%.2f\n",f1,f1)
	fmt.Println("%T,%.3f\n",f2,f2)

	var f3 = 2.55
	//注： Printf ； 输出类型
	fmt.Printf("%T\n",f3)
}