package main

type Musket struct {
	Gun
}

func newMusket() iGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket",
			power: 111,
		},
	}
}
