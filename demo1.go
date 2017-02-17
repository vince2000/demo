package main

import "fmt"

func main() {
	//show(4)
	/*a,b:=fanhui(11,13)
	fmt.Println(b)
	fmt.Println(a)*/
	//liang(fanhui,23,78))
	/*a,_:=fanhui(16,14)
	fmt.Println(a)*/
	map1:=map[string]string{}
	map1["ka"]="ka"
	map1["ka1"]="ka1"
	map1["ka2"]="ka2"
	map1["ka3"]="ka3"
	for  r:=range  map1{
		fmt.Println(r,map1[r])
	}
	var  b int
	var  a *int
	b=10;
	a=&b
	fmt.Println(b,a)
}
func show(a int){
	if a>10{
		fmt.Println("大于10")
	}else if a>5&&a<10{
		fmt.Println("5-10")
	}else{
		fmt.Println("zuixiao")
	}
}//weishane
func fanhui(a int,b int) (int,string){
	if a>=b {
		a1:=[2]int{}
		a1[0]=11
		a1[1]=12
		for _,r:=range a1{
			fmt.Println(r)
		}
		return a,"zhenque"
	}else {
		return b,"cuowu"
	}
}
/*
func liang(fanhui func()){
       a,_:=fanhui(23,78)
	fmt.Println(a)
}*/
