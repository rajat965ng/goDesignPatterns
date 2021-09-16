package main


type iBuilder interface {
	setWidowType()
	setFloorCount()
	getHouse() House
}

func getBuilder(name string) iBuilder  {
	if name == "normal" {
		return newNormalBuilder()
	}
	if name == "igloo" {
		return newIglooBuilder()
	}
	return nil
}