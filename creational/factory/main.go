package main

import "fmt"

func main()  {

	ak47,_ := getFactory("ak47")
	musket,_ := getFactory("musket")

	printVal(ak47)
	printVal(musket)
}

func printVal(gun iGun)  {
	fmt.Println(gun.getName())
	fmt.Println(gun.getPowder())
}