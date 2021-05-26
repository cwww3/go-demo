package main

import "fmt"

type Com interface {
	Operate()
}
type Com1 struct {
}

func (*Com1) Operate() {
	fmt.Println("操作")
}

type Der interface {
	Com
	Do()
}

type D1 struct {
	c Com
}

func (d *D1) Operate() {
	d.c.Operate()
}

func (d *D1) Do() {
	fmt.Println("装饰")
}

func main() {
	var d Der
	var c Com
	c = &Com1{}
	d = &D1{
		c: c,
	}
	d.Do()
	d.Operate()
}
