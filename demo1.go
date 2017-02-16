package main

import "fmt"

func main() {
	show(4)
}
func show(a int){
	if a>10{
		fmt.Println("大于10")
	}else if a>5&&a<10{
		fmt.Println("5-10")
	}else{
		fmt.Println("zuixiao")
	}
}