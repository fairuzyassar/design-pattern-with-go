package main

import "fmt"

type Category struct {
	name string
}

func (c *Category) getName() string {
	return c.name
}

type CategoryComposite struct {
	name string
	listOfCategory []Category
}

func (cc *CategoryComposite) getName() string {
	return cc.name
}

func (cc *CategoryComposite) addSubCategory(newSubCategory Category){
	cc.listOfCategory = append(cc.listOfCategory, newSubCategory)
}

func (cc *CategoryComposite) printListOfCategory(){
	fmt.Println(cc.getName())
	for _, category := range cc.listOfCategory{
		fmt.Println(category.getName())
	} 
}

func main(){
	smartphone := &CategoryComposite{name:"Smartphone"}
	smartphone.addSubCategory(Category{"Android"})
	smartphone.addSubCategory(Category{"Ios"})
	
	computer := &CategoryComposite{name:"Computer"}
	computer.addSubCategory(Category{"PC"})
	computer.addSubCategory(Category{"Notebook"})

	smartphone.printListOfCategory()
	computer.printListOfCategory()

}

