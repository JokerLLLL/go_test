package main

import "fmt"

func main() {
	var slice1 []int //zero value for slice is nil
	for i:= 0;i < 50 ;i ++ {
		fmt.Printf("slice len: %d,slice cap %d \n",len(slice1),cap(slice1))
		slice1 = append(slice1,i*2)
	}
   fmt.Println(slice1)

	slice2 := []int{0,1,2,3,4}
	slice3 := make([]int,5,7)
	slice4 := make([]int,10,10)
	for index := range slice4 {
		slice4[index] = index
	}
	// slice copy
	copy(slice3,slice4)
	fmt.Printf("value %v slice len: %d,slice cap %d \n",slice2,len(slice2),cap(slice2))
	fmt.Printf("value %v slice len: %d,slice cap %d \n",slice3,len(slice3),cap(slice3))
	fmt.Printf("value %v slice len: %d,slice cap %d \n",slice4,len(slice4),cap(slice4))

	//delete slice index = 8 from slice4
	slice4 = append(slice4[:8],slice4[9:]...)
	fmt.Println(slice4)

	// pop from front ,pop from back
	pop1 := slice4[0]
	slice4 = slice4[1:]
	pop2 := slice4[len(slice4)-1]
	slice4 = slice4[:len(slice4)-1]
	fmt.Println(pop1,pop2,slice4)
}
