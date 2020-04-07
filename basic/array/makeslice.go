package main

import (
	"fmt"
)

func main() {
	makeSlice()
	appendSlice()
	iteratorSlice()
}

func makeSlice() {
	a := make([]int, 5)
	printSlice1("a", a)

	b := make([]int, 0, 5)
	printSlice1("b", b)

	c := b[:2]
	printSlice1("c", c)

	d := c[2:5]
	printSlice1("d", d)

	appendSlice()
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice2(name string, s []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", name, len(s), cap(s), s)
}

func appendSlice() {
	var s []int
	printSlice2("first", s)

	s = append(s, 0)
	printSlice2("append 0", s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice2("append 1", s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice2("append more", s)

}

func iteratorSlice(){
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}


	for index ,value :=range pow{
		fmt.Printf("current index is %d, value is %d \n",index,value)
	}
}
