package main

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(name string) (iSportsFactory, error) {

	if name == "nike" {
		return &nike{}, nil
	}
	if name == "adidas" {
		return &adidas{}, nil
	}

	return nil, fmt.Errorf("Invalid factory name: %s\n", name)
}
