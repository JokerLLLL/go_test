//函数
package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func div(a,b int) (q int,r int) {
	q = a / b
	r = a % b
	return q, r
}

func sum(number ...int) (sum int,i int) {
	i = 0
	sum = 0
	fmt.Println(len(number))
	for index := range number {
		if i < index {
			i = index
		}
		sum += number[index]
	}
	return sum,i
}

//函数式编程
func apply(op func(a,b int) int,a,b int) int {
	//反射出该函的指针
	p := reflect.ValueOf(op).Pointer()
	func_name := runtime.FuncForPC(p).Name()
	fmt.Println(func_name)
	return op(a,b)
}

func main() {
	a,_ := div(13,4)
	fmt.Println(a)

	//sum := apply(func(a, b int) int {
	//	//	return a + b
	//	//},14,15)
	//	//fmt.Println(sum)
	fmt.Println(sum(1,2,45,5,5))
}
