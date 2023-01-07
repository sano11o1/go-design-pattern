package main

type adidas struct {
}

func (a *adidas) makeShirt(size int) IShirt {
	return &adidasShirt{
		Shirt: Shirt{
			size: size,
			logo: "adidas",
		},
	}
}

func (a *adidas) makeShoe(size int) IShoe {
	return &adidasShoe{
		Shoe: Shoe{
			size: size,
			logo: "アディダス",
		},
	}
}
