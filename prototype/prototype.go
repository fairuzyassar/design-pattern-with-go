package main

import "fmt"

type MPV interface{
	setName(newName string)
	setBrand(newBrand string) 
	ride() string
	clone() MPV
}

type Car struct {
	name string
	brand string
}

func (car *Car) setName(newName string){
	car.name = newName
}

func (car *Car) setBrand(newBrand string){
	car.brand = newBrand
}

func (car *Car) ride() string {
	return "I'm riding " + car.name + "-" +car.brand
}

func (car *Car) clone() MPV {
	return &Car{car.name, car.brand}
}

func main(){
	avanza := &Car{name: "avanza", brand: "toyota"}
	veloz := avanza.clone()
	veloz.setName("veloz")

	fmt.Println(avanza.ride())
	fmt.Println(veloz.ride())
}