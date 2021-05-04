package main

import "fmt"

// 只能传 [5]arr 的值
func printArr(arr [5]int)  {
	fmt.Println(arr)
}

//指针
func changeArr(arr *[5]int)  {
	arr[0] = 1000
}

//切片 slice
// 切片的第二个值不会包含在切片中，所以索引可以最索引+1
func sliceArr()  {
	arr := [...]int{0,11,22,33,44,55,66}
	s1 := arr[2:3] //半开半闭 切出 索引 2 的视图
	s2 := arr[:6]
	s3 := arr[3:]
	s4 := arr[2:3]  //  [22]
	s4 = s4[2:5]	// [44,55,66]
	//changeSlice(s1,s2,s3,s4)
	fmt.Println(s1,s2,s3,s4)

	fmt.Printf("s1=%v, len=%d, cap=%d\n",s1,len(s1),cap(s1))
	fmt.Printf("s2=%v, len=%d, cap=%d\n",s2,len(s2),cap(s2))

}
// []int 是一个数字的切片（视图）
func changeSlice(slice ...[]int)  {
	for i:=0;i < len(slice) ;i++  {
		slice[i][0] = 999
	}
}

func main() {
	//var arr1 [5]int
	//arr2 := [3]int{11,22,33}
	//arr3 := [...]int{1,3,5,7,9}
	//grid := [4][5]int{}
	//fmt.Println(arr1,arr2,arr3,grid)
	//for i:= 0 ;i <len(arr1);i++ {
	//	fmt.Println(arr1[i])
	//}
	//for _,value := range grid {
	//	fmt.Println(value)
	//	for index2,value2 := range value {
	//		fmt.Println(index2,value2)
	//	}
	//}
	//printArr(arr1)
	//changeArr(&arr1)
	//fmt.Println(arr1)

	//sliceArr()

	var arr = [7]int{0,1,2,3,4,5,6}
	s1 := arr[4:6]
	s2 := append(s1,991)
	s3 := append(s2,992)
	s4 := append(s3,994)
	fmt.Println(s2,s3,s4)
	// 超过 cap s3 s4 no longer view arr
	// 只会改变 s1 s2  ，，  s3,s4不在受到影响 //而且新生成的 arr len()长度会以 +1 增加，而cap()长度 x2 递增
	arr[5] = 555
	arr[6] = 666
	s4[2] = 191 // 说明 s4 view 同一个 arr
	fmt.Println(arr,s1,s2,s3,s4)

}
