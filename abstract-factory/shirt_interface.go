package main

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

/*
	AddidasShirt, NikeShirtに対して、インターフェースを実装するパターンもあり得る。
	https://refactoring.guru/ja/design-patterns/abstract-factory の疑似コードでは各具象に対してメソッドを生やしている。
	AddidasShirt, NikeShirtにShirtを埋め込むことで、AddidasShirt, NikeShirtはIShirtインターフェースを満たし、コードの重複を解消している。
	この共通化が正しいのかは、この演習では検討しない。
*/
type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}
