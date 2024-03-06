//Question 1: Create a struct named Rectangle with two fields, length and width, both of type int. Then, define a method named Area for the Rectangle struct that returns the area of the rectangle.

/*package main

import (
	"fmt"
)

type Rectangle struct {
	length int
	width  int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	rect := Rectangle{length: 5, width: 10}
	area := rect.Area()
	fmt.Println("Area of the rectangle:", area)
}*/

//Question 2: Define a struct called Circle with a field named radius of type float64. Write a method for the Circle struct named Perimeter that calculates and returns the perimeter of the circle. (Use math.Pi for the value of π).
/*
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	circle := Circle{radius: 3.0}
	perimeter := circle.Perimeter()
	fmt.Println("Perimeter of the circle:", perimeter)
}*/

//Question 3: Create a struct named Student with two fields: name of type string and grades of type []float64. Define a method for the Student struct named Average that calculates and returns the average of the grades.

/*package main

import (
	"fmt"
)

type Student struct {
	name   string
	grades []float64
}

func (s Student) Average() float64 {
	sum := 0.0
	for _, grade := range s.grades {
		sum += grade
	}
	return sum / float64(len(s.grades))
}

func main() {
	student := Student{name: "Alice", grades: []float64{85.0, 90.0, 88.0, 92.0}}
	average := student.Average()
	fmt.Printf("%s's average grade: %.2f\n", student.name, average)
}*/

//Question 4: Suppose you have a struct Counter with a field value of type int. Write a method named Increment that uses a pointer receiver to increment the value of the Counter by 1.
/*
package main

import (
	"fmt"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func main() {
	counter := Counter{value: 0}
	fmt.Println("Initial value of counter:", counter.value)
	counter.Increment()
	fmt.Println("Value of counter after incrementing:", counter.value)
}
*/

//Question 5: Create a struct named Person with two fields: firstName and lastName, both of type string. Define a method for the Person struct named FullName that returns the full name of the person (first name followed by last name, separated by a space).

package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func (p Person) FullName() string {
	return p.firstName + " " + p.lastName
}

func main() {
	person := Person{firstName: "John", lastName: "Doe"}
	fullName := person.FullName()
	fmt.Println("Full name:", fullName)
}
