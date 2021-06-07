//条件 循环
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)
func eval(a int,b int,p string) int {
	var result int
	/*
	switch p {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("非法字符串")
	}
	 */
	switch {
	case p == "*":
		result = a * b
	default:
		panic("die in illegal operation " + string(p))
	}
	return result
}

func loop(){

}
//十进制转二进制字符串
func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ;n > 0 ;n /= 2 {
		lsb := strconv.Itoa(n%2)
		result = lsb + result
	}
	return result
}

//单行读取文件
func printFile(filename string) {
	file, err := os.Open(filename)
	if(err != nil) {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(r io.Reader)  {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//读取整个文件
func readFile() {
	const filename = "abcc.txt"
	content,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",content)
	}
	// if条件可以赋值 但是 if里的局部变量
	if content,err := ioutil.ReadFile("abc.txt"); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(content)
	}
	fmt.Println(content)
}


func main() {
	readFile()
	fmt.Println(eval(2,3,"*"))

	fmt.Println(convertToBin(0),
		convertToBin(111232),
		convertToBin(343),
		convertToBin(999),
		)

	printFile("D:\\demo\\go\\src\\go_test\\abc.txt")

	s := `uu
er""eqrq
ffff
cccc
`
	printFileContents(strings.NewReader(s))
}
