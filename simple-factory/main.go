package main

import "fmt"

func main() {
	// クライアントコード
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	printDetails(ak47)
	printDetails(musket)
}

// ak47とmusketでinterfaceを統一したことにより、クライアントのコードではinterfaceを用いてやりとりする
func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
