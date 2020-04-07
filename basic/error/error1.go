package main

import (
	"fmt"
	"time"
)

/* type error interface {
	   Error() string
	}
   go 内建的error  error 是个接口，有实现error接口，就是error对象

*/
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string{
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error{
	return &MyError{time.Now(),"It didn't work"}
}

func main() {
	if err:=run(); err !=nil{
		fmt.Println(err)
		fmt.Printf("(%v, %T)\n", err, err)
	}
}