package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Printf("%s\n\n", t.S)
}

type F float64

func (f F) M() {
	fmt.Printf("%f\n\n", f)
}

type I2 interface {
	M()
}

func basics() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{4, 1}

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())

	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = F(math.SqrtPi)
	describe(i)
	i.M()

	var i2 I2
	describe(i2)
	//i2.M() // panic: runtime error: invalid memory address or nil pointer dereference
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
