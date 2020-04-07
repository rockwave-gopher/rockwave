package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

//func main() {main
func main() {

	const json1 = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
	age,_ := gojsonq.New().FromString(json1).From("age").GetR()
	fmt.Println(age)

}
