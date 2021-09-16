package main

import "fmt"

func getFactory(name string) (iGun,error) {

	if name=="ak47" {
		return newAk47(),nil
	}
	if name=="musket" {
		return newMusket(),nil
	}
	return nil,fmt.Errorf("Invalid input: %s",name)
}