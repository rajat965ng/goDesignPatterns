package main

import "fmt"

func main()  {
	iglooBuilder := getBuilder("igloo")
	iglooBuilder.setFloorCount()
	iglooBuilder.setWidowType()
	fmt.Println(iglooBuilder.getHouse())


	normalBuilder := getBuilder("normal")
	normalBuilder.setFloorCount()
	normalBuilder.setWidowType()
	fmt.Println(normalBuilder.getHouse())
}