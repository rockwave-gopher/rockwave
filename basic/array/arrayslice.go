package main

import "fmt"

func main(){
	var array1 = [6]int{1,2,3,4,5,6}
	fmt.Println(array1)

	a:=array1[0:2]
	b:=array1[1:3]
	fmt.Println(a,b)

	b[0]=100
	fmt.Println(a, b)
	fmt.Println(array1)
}
