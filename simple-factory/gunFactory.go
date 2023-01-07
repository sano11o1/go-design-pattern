package main

import "fmt"

func getGun(guntype string) (iGun, error) {
	if guntype == "ak47" {
		return newAK47(), nil
	}

	if guntype == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
