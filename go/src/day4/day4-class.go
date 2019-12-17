package day4

import (
	"fmt"
	"io"
)

type person struct {
	age int
}

func NewPerson(initialAge int, stdout io.Writer) person {
	//Add some more code to run some checks on initialAge
	p := person{age: initialAge}

	if initialAge < 0 {
		fmt.Fprintln(stdout, "Age is not valid, setting age to 0.")
		p.age = 0
	}

	return p
}

func (p person) amIOld(stdout io.Writer) {
	//Do some computation in here and print out the correct statement to the console
	switch {
	case p.age < 13:
		fmt.Fprintln(stdout, "You are young.")
	case p.age >= 13 && p.age < 18:
		fmt.Fprintln(stdout, "You are a teenager.")
	default:
		fmt.Fprintln(stdout, "You are old.")
	}
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age++
	return p
}

func run(stdin io.Reader, stdout io.Writer) {
	var T, age int

	fmt.Fscan(stdin, &T)

	for i := 0; i < T; i++ {
		fmt.Fscan(stdin, &age)

		p := NewPerson(age, stdout)
		p.amIOld(stdout)
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld(stdout)
		fmt.Fprintln(stdout)
	}
}
