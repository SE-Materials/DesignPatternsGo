package main

import "fmt"

type Password struct {
	value string
}

type Validator interface {
	Add(validator Validator)
	Validate(p Password)
}

type BaseValidator struct {
	password Password
	next     Validator
}

func (b *BaseValidator) Add(validator Validator) {
	if b.next != nil {
		b.next.Add(validator)
	} else {
		b.next = validator
	}
}

func (b *BaseValidator) Validate(p Password) {
	if b.next != nil {
		b.next.Validate(p)
	} else {
		fmt.Println("Base validator reached")
	}
}

// Empty validator
type EmptyValidator struct {
	BaseValidator
}

func (b *EmptyValidator) Validate(p Password) {
	if len(p.value) == 0 {
		fmt.Println("Empty password!")
	} else {
		b.BaseValidator.Validate(p)
	}
}

// Length validator
type LengthValidator struct {
	BaseValidator
}

func (b *LengthValidator) Validate(p Password) {
	if len(p.value) < 8 {
		fmt.Println("Min password length", 8)
	} else {
		b.BaseValidator.Validate(p)
	}
}

func main() {
	password := Password{value: "12345678"}

	rootValidator := EmptyValidator{}
	lengthValidator := LengthValidator{}
	rootValidator.Add(&lengthValidator)

	rootValidator.Validate(password)
}
