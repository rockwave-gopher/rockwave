package main

import (
	"github.com/davecgh/go-spew/spew"
	"errors"
)


func main() {
	spew.Config.Indent = "\t"
	i:=0
	s:="哈哈"
	m:=map[int]string{1:"1",2:"2"}
	e:=errors.New("嘿嘿，错误")
	p:=person{Name:"张三"}
	spew.Dump(spew.Config)
	spew.Dump(i,s,m,e,p)
}

type person struct {
	Name string
}