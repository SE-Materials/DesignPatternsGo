package main

import "fmt"

type Person struct {
	// Address
	StreetAddress, PostCode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

// --- Person builder
type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

// --- Person address builder
type PersonAddressBuilder struct {
	PersonBuilder
}

func (pb *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	pb.person.StreetAddress = address
	return pb
}

func (pb *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pb.person.City = city
	return pb
}

func (pb *PersonAddressBuilder) WithPostCode(postcode string) *PersonAddressBuilder {
	pb.person.PostCode = postcode
	return pb
}

// --- Person job builder

type PersonJobBuilder struct {
	PersonBuilder
}

func (pj *PersonJobBuilder) At(company string) *PersonJobBuilder {
	pj.person.CompanyName = company
	return pj
}

func (pj *PersonJobBuilder) AsA(title string) *PersonJobBuilder {
	pj.person.Position = title
	return pj
}

func (pj *PersonJobBuilder) Earning(amount int) *PersonJobBuilder {
	pj.person.AnnualIncome = amount
	return pj
}

// --- Get facets

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

// --- main
func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostCode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000)

	person := pb.Build()
	fmt.Println(person)

}
