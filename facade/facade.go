package main

import "fmt"

type Warehousing struct {}

func (w *Warehousing) updateStock() string {
	return "Reduce stock by one"
}

type Packaging struct {}

func (p *Packaging) goodsPackaging() string {
	return "Package good using carton box"
}

type Delivery struct {}

func (d *Delivery) goodsDelivery() string {
	return "Deliver using truck"
}


type ShopFacade struct {
	warehouse *Warehousing
	packaging *Packaging
	delivery *Delivery
}

func customerOrder() *ShopFacade{
	return &ShopFacade{&Warehousing{}, &Packaging{}, &Delivery{}}
}

func (sf *ShopFacade) orderProcessing() {
	fmt.Println(sf.warehouse.updateStock())
	fmt.Println(sf.packaging.goodsPackaging())
	fmt.Println(sf.delivery.goodsDelivery())
}

func main(){
	shop := customerOrder()
	shop.orderProcessing()
}