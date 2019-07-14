package main

import (
	"fmt"
)

func main() {

	student := Student{"lzr", 0, 0, Result{0,0}}
	student.math, student.eng = student.result.add(student.math,student.eng)
	fmt.Println(student)
}

type Student struct {
	name string
	eng int
	math int
	result Adder
}
type Adder interface {
	add(eng, math int)(int, int)
}
type Result struct {
	eng int
	math int
}

func (r Result) add(eng, math int)(int, int)  {
	r.math += 100
	r.eng += 100
	return r.math, r.eng
}


