package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes慕课网！!" // len 获取到 字符串编码的长度 UTF-8
	fmt.Println(len(s))
	fmt.Println()
	fmt.Printf("%s \n",s)
	fmt.Printf("%x \n",s)
	fmt.Printf("%s \n",[]byte(s))
	fmt.Printf("%X \n",[]byte(s))

	for i,ch := range s { //ch is rune
		fmt.Printf("(%d %X)",i,ch)
	}
	fmt.Println()
	fmt.Println("run count:", utf8.RuneCountInString(s))

	//通过获取byte 获取rune
	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes) // 通过 bytes 取出 rune 和 size
		bytes = bytes[size:]
		fmt.Printf("(%c %d %X ) ",r,r,r)
	}
	fmt.Println()
	runes := []rune(s) //重新开辟空间 编译 rune
	for i,v := range s { //单个字直接遍历
		fmt.Printf("(%d,%c) ",i,v)
	}
	fmt.Println()
	for i,v := range runes {
		fmt.Printf("(%d,%c) ",i,v)
	}

}
