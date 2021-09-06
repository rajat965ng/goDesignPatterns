package main

type nike struct {
}

func (*nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe{
			size: 10,
			logo: "adidas",
		},
	}
}

func (*nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt{
			size: 32,
			logo: "adidas",
		},
	}
}
