package main

import "fmt"


type Point struct{
	X float64
	Y float64
}

func createPoint(X float64, Y float64) Point {
	point := Point {X: X, Y: Y}

	return point
}

func printPoint(point Point) {
	fmt.Printf("Point(X: %.2f, Y: %.2f)\n", point.X, point.Y)
}

func main(){
	p := createPoint(3.5, 5.3)
	printPoint(p)
}