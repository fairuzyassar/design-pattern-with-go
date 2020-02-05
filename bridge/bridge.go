package main

import "fmt"

type Shape interface {
	drawing() string
}

type Circle struct {}

func (c *Circle) drawing() string {
	return "Drawing Circle"
}

type Square struct {}

func (c *Square) drawing() string {
	return "Drawing Square"
}

type Color interface {
	apply() string
}

type Blue struct {
	shape Shape
}

func (b *Blue) apply() string {
	return "Apply Blue Color Into " + b.shape.drawing()
}

type Red struct {
	shape Shape
}

func (r *Red) apply() string {
	return "Apply Red Color Into " + r.shape.drawing()
}

func main(){
	redCircle := &Red{&Circle{}}
	blueCircle := &Blue{&Circle{}}
	redSquare := &Red{&Square{}}
	blueSquare := &Red{&Square{}}

	fmt.Println(redCircle.apply())
	fmt.Println(blueCircle.apply())
	fmt.Println(redSquare.apply())
	fmt.Println(blueSquare.apply())

}