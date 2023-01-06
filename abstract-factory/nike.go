package main

type Nike struct {
}

func (a *Nike) makeShirt(size int) IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			size: size,
			logo: "NIKE",
		},
	}
}

func (a *Nike) makeShoe(size int) IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			size: size,
			logo: "NIKE",
		},
	}
}
