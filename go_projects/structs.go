package main

import "fmt"

type Coordinate struct {
	x,y int
}

type Rectangle struct {
	a Coordinate
	b Coordinate
}

func width(rect Rectangle) int {
	return (rect.b.x - rect.a.x)
}


func length(rect Rectangle)int {
	return (rect.a.y - rect.b.y);
}

func area(rect Rectangle) int {
	return length(rect) * width(rect)
}

func paramter(rect Rectangle) int { 
	return (width(rect)* 2)+ (length(rect) + 2)
}


func printInfo(rect Rectangle){
	fmt.Println("area ", area(rect));
	fmt.Println("paramater ,", paramter(rect));

}


func main(){
rect := Rectangle{a : Coordinate{0,7}, b: Coordinate{10,0}}


printInfo(rect)


rect.a.y *= 2
rect.b.x *= 2


printInfo(rect)


}


