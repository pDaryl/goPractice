package main

import (
	"fmt"
)

type Employee struct{
	Name string
	ID int 
	Address Address
}

type Address struct {
	Street string
	City string 
	PostalCode string 
}

func createEmployee (name string, id int, street string, city string, postalCode string) Employee{
	employee := Employee{Name: name, ID: id, Address: Address{Street: street, City: city, PostalCode: postalCode}}
	return employee
}

func printEmployee(employee Employee){
	fmt.Printf(
		`Employee(Name: %s, ID: %d, Address: [Street: %s, City: %s, PostalCode: %s])
`,
		employee.Name,
		employee.ID,
		employee.Address.Street,
		employee.Address.City,
		employee.Address.PostalCode,
	)
	// Employee(Name: <name>, ID: <id>, Address: [Street: <street>, City: <city>, PostalCode: <postalCode>])

}

func main(){
e:= createEmployee("Daryl", 100, "Zeta", "Belle Chasse", "70037")
printEmployee(e)
}