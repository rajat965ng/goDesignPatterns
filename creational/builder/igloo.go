package main


type igloo struct{
	windowType string
	floorCount int
}

func newIglooBuilder() *igloo  {
	return &igloo{}
}

func (i* igloo) setWidowType() {
	i.windowType = "iglooWindow"
}

func (i* igloo) setFloorCount() {
	i.floorCount = 1
}

func (i* igloo) getHouse() House {
	return House{
		floorCount:i.floorCount,
		windowType:i.windowType,
	}
}

