package main

import (
	"fmt"
	"math"
)

//结构体  Go 没有类。不过你可以为结构体类型定义方法。方法就是一类带特殊的 接收者 参数的函数。
type Vertex struct {
	X,Y float64
}

//方法就是一类带特殊的 接收者 参数的函数。在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者
func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//你可以为指针接收者声明方法,以指针为接收者的方法被调用时，接收者既能为值又能为指针：
//使用指针接收者的原因有二：
//首先，方法能够修改其接收者指向的值。
//其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}


func main(){
	fmt.Println(" 值接收者---------------------")
	v := Vertex{3,4}
	fmt.Println(v.Abs())


	fmt.Println("指针接收者---------------------")
	v.Scale(10)
	fmt.Println(v.Abs())



}



