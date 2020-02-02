package main

import "fmt"

type Transport interface{
	goToAirport() string
}

type Bus struct{}
type Taxi struct{}
type MRT struct{}

func (bus *Bus) goToAirport() string {
	return "Go to airport by bus will be charge $3"
}

func (taxi *Taxi) goToAirport() string {
	return "Go to airport by taxi will be charge $9"
}

func (mrt *MRT) goToAirport() string {
	return "Go to airport by MRT will be charge $7"
}

func TransportFactory(t string) Transport{
	switch t {
	case "Bus":
		return &Bus{}
	case "Taxi":
		return &Taxi{}
	case "MRT":
		return &MRT{}
	default:
		return nil
	}
}

func main(){
	bus := TransportFactory("Bus")
	taxi := TransportFactory("Taxi")
	mrt := TransportFactory("MRT")

	fmt.Println(bus.goToAirport())
	fmt.Println(taxi.goToAirport())
	fmt.Println(mrt.goToAirport())
}
