// Force users to use builder rather than the object directly!
package main

import "strings"

// --- Email
type email struct {
	From, To, Subject, Body string
}

// --- Email builder
type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("From should contain @")
	}

	b.email.From = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.To = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.Subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.Body = body
	return b
}

// --- send mail utility
func sendMailImpl(email *email) {

}

// ---
type build func(*EmailBuilder)

func SendEmail(action build) {
	eb := EmailBuilder{}
	action(&eb)
	sendMailImpl(&eb.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
