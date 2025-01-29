package main

import "fmt"

type Employees struct {
	EmployeeName string
	ID string
	Tasks []string
}


func (e *Employees) addTasks(task string) {
	e.Tasks = append(e.Tasks, task)
}

func listEmployees(e Employees) Employees {
employee := Employees{
	EmployeeName: e.EmployeeName,
	ID: e.ID,
	Tasks: e.Tasks,
}
return employee
}

func findMostTasks(employees []Employees) Employees {
mostTasks := 0
var topEmployee Employees

for _, employee := range employees{
	if len(employee.Tasks) > mostTasks{
		mostTasks = len(employee.Tasks)
		topEmployee = employee
	}
}
return topEmployee
}

func main(){

employees := []Employees{
	{
		"daryl", 
		"69", 
		[]string{"clean computer", "take a bath"},
	}, 
	{
		"Rae", 
		"1", 
		[]string{"eat a bagel", "read a book", "pick a movie"},
	}, 
	{
		"Boo", 
		"55", 
		[]string{"meow a lot", "cuddle with daryl"},
	}, 
}

employees[0].addTasks("go to gym")
employees[0].addTasks("clean the house")
employees[2].addTasks("use the litter box")

fmt.Println("All Employees:")
for _, e := range employees{
fmt.Println("Name:",e.EmployeeName, "ID:",e.ID)
fmt.Println("Tasks:")
for i, task := range e.Tasks{
	fmt.Printf("%d. %s\n", i+1, task)
}
}
fmt.Println()
fmt.Println("Top Employee:")
topEmployee := findMostTasks(employees)
fmt.Printf("Name: %s\n", topEmployee.EmployeeName)
fmt.Printf("ID: %s\n", topEmployee.ID)
fmt.Println("Tasks:")
for i, t := range topEmployee.Tasks{
fmt.Printf("  %d. %s\n", i+1, t)
}

}




