package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	person := &Person{Name: "Alice", Age: 25}
	fmt.Println(person.IsAdult())
}
