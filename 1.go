package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printPerson(pers Person) {
	fmt.Println("Name:", pers.name)
	pers.name = "John Smith"
	fmt.Println("AGe:", pers.age)
}
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

}
