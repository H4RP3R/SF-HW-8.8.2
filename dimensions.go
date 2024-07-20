package main

type dimensionsCm struct {
	length   float64
	width    float64
	height   float64
	unitType UnitType
}

func (d dimensionsCm) Length() Unit {
	return Unit{d.length, d.unitType}
}

func (d dimensionsCm) Width() Unit {
	return Unit{d.width, d.unitType}
}

func (d dimensionsCm) Height() Unit {
	return Unit{d.height, d.unitType}
}

func NewDimensionsCm(length, width, height float64) dimensionsCm {
	return dimensionsCm{length, width, height, CM}
}

type dimensionsInch struct {
	length   float64
	width    float64
	height   float64
	unitType UnitType
}

func (d dimensionsInch) Length() Unit {
	return Unit{d.length, d.unitType}
}

func (d dimensionsInch) Width() Unit {
	return Unit{d.width, d.unitType}
}

func (d dimensionsInch) Height() Unit {
	return Unit{d.height, d.unitType}
}

func NewDimensionsInch(length, width, height float64) dimensionsInch {
	return dimensionsInch{length, width, height, Inch}
}
