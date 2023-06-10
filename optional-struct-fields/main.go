package main

import "fmt"

type Foo struct {
	name string
}

type Bar struct {
	foo *Foo
}

func (f *Foo) GetName() string {
	return f.name
}

func (b *Bar) PrintLose() {
	fmt.Println("Lose")
}

func (b *Bar) PrintWin() {
	fmt.Printf("%s, you win", b.foo.GetName())
}

func main() {
	// b := Bar{foo: &Foo{name: "John"}}
	b := Bar{}

	for {
		n := 0
		fmt.Println("Enter a number: ")
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		b.PrintLose()
	}

	b.PrintWin()
}
