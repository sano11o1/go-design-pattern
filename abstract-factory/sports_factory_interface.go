package main

import "fmt"

type SportsFactory interface {
	makeShirt(size int) IShirt
	makeShoe(size int) IShoe
}

func GetSportsFactory(brand string) (SportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	} else if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
