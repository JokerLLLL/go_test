//指针
//go 都是值传递

package main

import "fmt"

func doString(p *string)  {
	*p = "8888888888888888"
}

func swap(a,b *int)  {
	*a,*b = *b,*a
}



func main() {

	var aa string = "9999999999999999999999999999999999999999"
	doString(&aa)
	fmt.Println(aa)
	var test1 = 111
	var test2 = 222
	fmt.Printf("test1: %d test2: %d \n",test1,test2)


	//swap1(&test1,&test2)
	//fmt.Printf("test1: %d test2: %d \n",test1,test2)

	test1,test2 = test2,test1

	fmt.Printf("test1: %d test2: %d \n",test1,test2)

}
