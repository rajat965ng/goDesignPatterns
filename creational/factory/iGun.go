package main

type iGun interface {
	setName(name string)
	setPowder(powder string)
	getName() string
	getPowder() string
}

type gun struct{
	name string
	powder string
}

func (g* gun) setName(name string) {
	g.name = name
}

func (g* gun) setPowder(powder string) {
	g.powder = powder
}

func (g* gun) getName() string {
	return g.name
}

func (g* gun) getPowder() string {
	return g.powder
}


