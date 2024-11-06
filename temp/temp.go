package temp

import (
	"fmt"
)

func TestFunc() {
	hr := 201
	mi := 3001

	a := hr / 24
	b := hr % 24

	x := mi / 60
	y := mi % 60
	fmt.Println("")
	fmt.Println("hours")
	fmt.Printf("hour %d", a)
	//fmt.Println(b)
	fmt.Printf("remainder minutes %d", b)

	fmt.Println(" \n minutes")
	fmt.Println(x)
	fmt.Println(x / 24)
	fmt.Println(y)
}

type Temp struct {
	age  int
	name string
}

func (t Temp) New() *Temp {
	return &Temp{}
}

func TestDiv() {
	n := 140
	i := n / 60
	r := n % 60
	fmt.Println(i)
	fmt.Println(r)
}

func Compare() {
	t := Temp{
		age:  2,
		name: "byorn",
	}

	t1 := Temp{
		age:  2,
		name: "byorn",
	}

	fmt.Println(t1 == t)

	t.New()
	t.age = 3
	fmt.Printf("t is %v", t)

	fmt.Printf("t1 is %v", t1)
}
