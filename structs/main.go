package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // Embedding a struct in a struct
}

type person2 struct {
	firstName   string
	lastName    string
	contactInfo // saving keystrokes by just declaring the struct
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	// Go structs uses order to assign the values if no parameters are passed
	// always use the parameters and colons to define which value is assigned to what
	jake := person{firstName: "Jake", lastName: "Chan"}

	// when calling var jake person, Go assigns zero values to all the fields in the struct
	// int zero value = 0
	// string zero value = ""
	// float zero value = 0
	// boolean zero value = false

	fmt.Println(jake.firstName, jake.lastName)

	// Printf("%+v") prints a much more verbose version of the variable
	fmt.Printf("%+v\n", jake)

	// updating the struct
	jake.firstName = "Jackie"

	fmt.Println(jake.firstName, jake.lastName)

	// semi colons are inserted automatically by the lexer, which is at the end of non-blank lines
	// for that reason, we need to close the curly braces at the string literal
	personWithContact := person{
		firstName: "test1",
		lastName:  "person",
		contact: contactInfo{
			email: "foobar@gmail.com",
			zip:   54321}}

	personWithContact2 := person2{
		firstName: "test2",
		lastName:  "person",
		contactInfo: contactInfo{
			email: "foobar@gmail.com",
			zip:   54321}}

	personWithContact.print()
	personWithContact2.print()

	// getting the address of the struct
	// personWithContact2Pointer := &personWithContact2
	// personWithContact2Pointer.updateName("hello")
	// the sentence below is a short cut to that code
	personWithContact2.updateName("hello")

	personWithContact2.print()
}

// setting up a receiver with a struct
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// setting up a receiver with a struct
func (p person2) print() {
	fmt.Printf("%+v\n", p)
}

// The receiver has to be a pointer to reference the actual struct
// by default, go passes by value, which creates a new copy of the value being passed in
// if we didnt pass in a struct, go will create a brand new person2 struct
// the orignal struct will remain unchanged
func (pointerToPerson *person2) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}
