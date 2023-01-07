package main

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

/*
	AddidasShoe, NikeShoeに対して、インターフェースを実装するパターンもあり得る。
	https://refactoring.guru/ja/design-patterns/abstract-factory の疑似コードでは各具象に対してメソッドを生やしている。
	AddidasShoe, NikeShoeにShoeを埋め込むことで、AddidasShoe, NikeShoeはIShoeインターフェースを満たし、コードの重複を解消している。
	この共通化が正しいのかは、この演習では検討しない。
*/

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}
