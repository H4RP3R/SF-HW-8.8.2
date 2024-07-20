package main

type autoBMW struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

func (a autoBMW) Brand() string {
	return a.brand
}

func (a autoBMW) Model() string {
	return a.model
}

func (a autoBMW) Dimensions() Dimensions {
	return a.dimensions
}

func (a autoBMW) MaxSpeed() int {
	return a.maxSpeed
}

func (a autoBMW) EnginePower() int {
	return a.enginePower
}

func NewAutoBMW(model string, maxSpeed, enginePower int, length, width, height float64) autoBMW {
	return autoBMW{"BMW", model, NewDimensionsCm(length, width, height), maxSpeed, enginePower}
}

type autoMercedes struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

func (a autoMercedes) Brand() string {
	return a.brand
}

func (a autoMercedes) Model() string {
	return a.model
}

func (a autoMercedes) Dimensions() Dimensions {
	return a.dimensions
}

func (a autoMercedes) MaxSpeed() int {
	return a.maxSpeed
}

func (a autoMercedes) EnginePower() int {
	return a.enginePower
}

func NewAutoMercedes(model string, maxSpeed, enginePower int, length, width, height float64) autoMercedes {
	return autoMercedes{"Mercedes", model, NewDimensionsCm(length, width, height), maxSpeed, enginePower}
}

type autoDodge struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

func (a autoDodge) Brand() string {
	return a.brand
}

func (a autoDodge) Model() string {
	return a.model
}

func (a autoDodge) Dimensions() Dimensions {
	return a.dimensions
}

func (a autoDodge) MaxSpeed() int {
	return a.maxSpeed
}

func (a autoDodge) EnginePower() int {
	return a.enginePower
}

func NewAutoDodge(model string, maxSpeed, enginePower int, length, width, height float64, unitType UnitType) autoDodge {
	var dimensions Dimensions
	switch unitType {
	case Inch:
		dimensions = NewDimensionsInch(length, width, height)
	case CM:
		dimensions = NewDimensionsCm(length, width, height)
	}

	return autoDodge{"Dodge", model, dimensions, maxSpeed, enginePower}
}
