package main

import "fmt"

func main() {
	adidasFacotry, _ := GetSportsFactory("adidas")
	nikeFacotry, _ := GetSportsFactory("nike")

	adidasShoe := adidasFacotry.makeShoe(26)
	nikeShoe := nikeFacotry.makeShoe(27)

	adidasShirt := adidasFacotry.makeShirt(160)
	nikeShirt := nikeFacotry.makeShirt(170)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

}

/*
	以下のコードAbstract Facotyを使うクライアントコードとする
	引数はinterfaceのため、addidasか、clientかを気にする必要はない
*/
func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
