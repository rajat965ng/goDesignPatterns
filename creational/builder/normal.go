package main

type normal struct{
	windowType string
	floorCount int
}

func newNormalBuilder() *normal  {
	return &normal{}
}

func (n* normal) setWidowType() {
	n.windowType = "normalWindow"
}

func (n* normal) setFloorCount() {
	n.floorCount = 2
}

func (n* normal) getHouse() House {
	return House{
		floorCount:n.floorCount,
		windowType:n.windowType,
	}
}

