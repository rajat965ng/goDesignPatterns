package main

type adidas struct {
}

func (*adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe{
			size: 10,
			logo: "adidas",
		},
	}
}

func (*adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt{
			size: 32,
			logo: "adidas",
		},
	}
}
