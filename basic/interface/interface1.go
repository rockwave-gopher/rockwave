package main

import (
	"fmt"
	"math"
)


//类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。
//隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。
//因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。
//接口也是值。它们可以像其它值一样传递。
//接口值可以用作函数的参数或返回值。
//在内部，接口值可以看做包含值和具体类型的元组： (value, type)
//接口值保存了一个具体底层类型的具体值。
//接口值调用方法时会执行其底层类型的同名方法。

type I interface {
	M()
}

type T struct {
	s string
	n string
}

type VertexNew struct {

	X,Y int
}

func (t *T) M(){
	fmt.Println(t.s+" "+t.n)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main(){
	var i I

	i = &T{"Hello","World"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()


	//指定了零个方法的接口值被称为 *空接口：*
	//interface{}
	//空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
	var zeroI interface{}
	describeI(zeroI)
	zeroI = 42
	describeI(zeroI)
	zeroI = VertexNew{3,4}
	describeI(zeroI)

	//断言 提供了访问接口值底层具体值的方式。
	v,ok:= zeroI.(VertexNew)
	fmt.Println(v,ok) //return {3,4} true

	f, ok := zeroI.(float64)
	fmt.Println(f,ok) //return  0 false



}

func describeI(i interface{}){
	fmt.Printf("(%v, %T)\n", i, i)

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}