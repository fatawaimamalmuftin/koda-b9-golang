package main

import (
	"fmt"
)

func main() {
	// manifest
	var hallo string 
	hallo = "hallo"
	fmt.Println(hallo)
	hallo = "word"
	fmt.Println(hallo)
	word:="word hello"
	fmt.Println(word)

	fmt.Println("hallo")
	fmt.Println("word")
	//aritmatika-------------------------

	var a int8 = 1
	var b uint8 = 2
	c:= uint8(a) + b

	// kondisi if
	if c > 5 {
		fmt.Println("hore")
	}else{
		fmt.Println("reho")
	}

	// kondisi swich
	switch true {
		case c > 5:
			fmt.Println("hore")
		case c > 2:
			fmt.Println("reho")
		default:
			fmt.Println("error")
	}
}
