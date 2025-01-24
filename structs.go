package main

import (
	"fmt"
)


type Rectangle struct{
	Length float64
	Width float64
}

func createRectangle (Length float64, Width float64) Rectangle {
	rectangle := Rectangle{Length: Length, Width: Width}
	return rectangle
}

func printRectangle (rectangle Rectangle){
	fmt.Printf("Rectangle(Length: %.1f, Width: %.1f)\n", rectangle.Length, rectangle.Width)
}

func main(){
	rectangle := createRectangle(4.5, 5.4)
	printRectangle(rectangle)
}