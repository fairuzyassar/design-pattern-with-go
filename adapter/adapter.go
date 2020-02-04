package main

import "fmt"

type PhoneCharger struct {}

func (p *PhoneCharger) chargingBattery() string {
	return "Charging using Type A plugin"
}

type Converter interface {
	charging() string
}

type Adapter struct {
	charger *PhoneCharger
}

func phoneAdapter() *Adapter{
	return &Adapter{&PhoneCharger{}}
}

func (a *Adapter) charging() string{
	fmt.Println("Charging using Type C Converter")
	return a.charger.chargingBattery()

}

func main(){
	adapter := phoneAdapter()
	fmt.Println(adapter.charging())
}