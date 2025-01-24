package main

import "fmt"


type Person struct {
	name string
	age int 
}

func createPerson(name string, age int) Person {
p1 := Person {name: name , age: age}
return p1

}
func main(){
person := createPerson("Daryl", 28)

fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
}