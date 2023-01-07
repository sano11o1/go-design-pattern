package main

type AK47 struct {
	Gun
}

func newAK47() iGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47",
			power: 123,
		},
	}
}
