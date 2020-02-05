package main

import "fmt"

type HouseTemplate interface {
	floorNumber() int
	bedroomNumber() int
}

type HousingAbstract struct {
	template HouseTemplate
}

func (h *HousingAbstract) build() {
	fmt.Println("\nBuild House with :")
	fmt.Printf("%d Floor\n", h.template.floorNumber())
	fmt.Printf("%d Bedroom\n", h.template.bedroomNumber())
}

type SmallHouse struct {
	*HousingAbstract
}

func (sh *SmallHouse) floorNumber() int {
	return 1
}

func(sh *SmallHouse) bedroomNumber() int {
	return 1
}

type MediumHouse struct {
	*HousingAbstract
}

func (mh *MediumHouse) floorNumber() int {
	return 2
}

func(mh *MediumHouse) bedroomNumber() int {
	return 2
}

type BigHouse struct {
	*HousingAbstract
}

func (bh *BigHouse) floorNumber() int {
	return 3
}

func(bh *BigHouse) bedroomNumber() int {
	return 3
}

func main() {
	smallHouse := &HousingAbstract{&SmallHouse{}}
	smallHouse.build()

	mediumHouse := &HousingAbstract{&MediumHouse{}}
	mediumHouse.build()

	bigHouse := &HousingAbstract{&BigHouse{}}
	bigHouse.build()

}

