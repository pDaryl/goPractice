package main

import (
	"fmt"
)


type Circle struct{
	Radius float64
	Color string
}

func createCircle(radius float64, color string) Circle {
	circle := Circle{Radius: radius, Color: color}
	return circle
} 

func printCircle(circle Circle) {
	fmt.Printf("Circle(Radius: %.2f, Color: %s)\n", circle.Radius, circle.Color)
} 

func main(){
c := createCircle(7.25, "green")
printCircle(c)
}