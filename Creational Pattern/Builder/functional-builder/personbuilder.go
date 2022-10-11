package main

import "fmt"

type Person struct {
	Name, WorksAt string
}

type Action func(*Person)
type PersonBuilder struct {
	person  Person
	actions []Action
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.Name = name
	})
	return b
}

func (b *PersonBuilder) WorksAt(at string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.WorksAt = at
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, act := range b.actions {
		act(&p)
	}

	return &p
}

func main() {
	pb := PersonBuilder{}
	person := pb.Called("Aniket").
		WorksAt("Siemens").
		Build()

	fmt.Println(*person)
}
