package main

import "fmt"

func main() {

	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	adidasShirt := adidasFactory.makeShirt()
	adidasShoe := adidasFactory.makeShoe()

	nikeShirt := nikeFactory.makeShirt()
	nikeShoe := nikeFactory.makeShoe()

	fmt.Println(nikeShirt)
	fmt.Println(nikeShoe)

	fmt.Println(adidasShoe)
	fmt.Println(adidasShirt)

}
