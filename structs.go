package main

import "fmt"

type Student struct{
	Name string
	Grades []int
	ID string
}

func listStudents(students []Student){
	for _, student := range students{
		fmt.Printf("Student: %s, ID: %s, Grades: %v\n", student.Name, student.ID, student.Grades)
	}
}

func (s *Student) averageGrades() float32 {
	sum := 0
	for _, g := range s.Grades{
		sum += g
	}
	 GPA := float32(sum) / float32(len(s.Grades))

	return GPA
}

func getTopStudent(students []Student) Student {
	var topStudent Student
	var highestGPA float32

	for _, student := range students {
		currentGPA := student.averageGrades()
		if currentGPA > highestGPA {
			highestGPA = currentGPA
			topStudent = student
		}
	}
	return topStudent
}

func (s *Student) printStudentDetails() {
	fmt.Printf("Student: %s, ID: %s, Grades: %v, GPA: %.2f\n", s.Name, s.ID, s.Grades, s.averageGrades())
}


func main(){

students := []Student{{
	Name:   "Daryl",
	Grades: []int{80, 78, 92, 88},
	ID:     "55",
}, 
{
	Name:   "Rae",
	Grades: []int{82, 79, 90, 89},
	ID:     "69",
}, 
{
	Name:   "Boo",
	Grades: []int{99, 90, 80, 70},
	ID:     "1",
},
}

fmt.Println("All Students:")
for _, s := range students {
	s.printStudentDetails()
}

fmt.Println("Top Student Details:")
topStudent := getTopStudent(students)
topStudent.printStudentDetails()

}




