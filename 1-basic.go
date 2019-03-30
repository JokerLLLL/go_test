package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//包内变量
var aa = 1223
var ss string = "xxx"
var(
	bb = "hello world"
)

func variable()  {
	//函数变量
	var  a int
	var  b string
	fmt.Println(a,b)
}
func variableZeroValue()  {
		//默认0
		var a int
		//空字符串
		var s string = ""
		// %s打印有效字段  %q 可以打印空字符串加上引号
		fmt.Printf("%d %s %q\n",a,s,s)
}
func variableType()  {
		a,b,c,d,e := 1,2,3,"true",true
		println(a,b,c,d,e)
}

func Euler(){
	//复数
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(
		cmplx.Pow(math.E,1i*math.Pi)+1)
	//3位数 float
	fmt.Printf("%.3f",
		cmplx.Exp(1i*math.Pi)+1)
}
//强制类型转换
func triangle()	{
	a,b := 3,4
	var c int
	c = int(math.Sqrt(float64(a*a)+float64(b*b)))
	fmt.Println(c)
}
// 常量 枚举
func consts()  {
	const(
		filename = "abc.txt"
		aa = 1
		bb = 5
		cpp = iota
		php = iota
		java = iota
		)
	const(
	//b kb mb tb
		b = 1<<(10*iota)
		kb
		mb
		gb
		tb
	)
	var c int
	c = int(math.Sqrt(aa+bb))
	fmt.Println(c)
	fmt.Println(b,kb,mb,gb,tb)
}
func main() {
	//fmt.Println("hello world")
	//variable()
	//variableZeroValue()
	//variableType()
	//fmt.Println(aa,bb,ss);
	//Euler()
	//triangle()
	consts()
}
