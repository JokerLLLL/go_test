package main

import "fmt"

//寻找最长含有重复字符的子串
//abcabcbb -> abc 3
//bbbbb -> b 1
//pwwkew -> wke 3
// abc -> abc 3
// "" -> ""  0


func searchLengthSonString(s string) (length int) {
	var max_lenght  int
	var start int

	lastCurrent := map[byte]int{}
	// ch 是 ascii码
	for i,ch := range []byte(s) {
		if last_index,ok := lastCurrent[ch]; ok && last_index >= start {
			if max_lenght < (i - start) {
				max_lenght = i - start
			}
			start = last_index + 1
		}
		//更新 map last位置
		lastCurrent[ch] = i
	}
	//最后一次判断
	if max_lenght < (len(s) - start) {
		max_lenght = len(s) - start
	}

	return max_lenght
}

func main() {
	map1 := map[string]string {
		"java":"19k",
		"go":"15k",
		"php":"8k",
		"js":"13k",
	}
	map2 := make(map[string]int)	//map2 === empty map
	var map3 map[string]int // map3 == nil
	if map2 == nil {
		fmt.Println(map1,map2,map3)
	}
	//map 遍历 遍历的时候 map 是无序的 （hash map）
	for k,v := range map1 {
		fmt.Printf("%v, %v \n",k,v)
	}
	//key不存在 会获得go 的 zero value
	fmt.Print(map1["php"],map1["c#"])
	if value,ok := map1["c#"] ; ok {
		fmt.Printf(value)
	}else{
		//panic("key does not exists")
	}

	//删除map
	delete(map1,"php")
	fmt.Println(map1)

	fmt.Println(searchLengthSonString(""))
}
