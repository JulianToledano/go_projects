package main

import "fmt"

func main() {
	p := new(pile)
	a := element{content: 1}
	b := element{content: 1}
	c := element{content: 1}
	d := element{content: 1}
	p.push(a)
	p.push(b)
	p.push(c)
	p.push(d)

	fmt.Println(p)
}
