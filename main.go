package main

import (
	"fmt"
	"math"
	"strings"
)

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		switch t {
		case CM:
			return math.Round(value*2.54*10) / 10
		case Inch:
			return math.Round(value*0.3937*10) / 10
		}
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

func main() {
	// Mercedes
	mercedes := NewAutoMercedes("AMG SL 43", 250, 381, 471.9, 189.0, 132.0)
	fmt.Println("Brand:", mercedes.Brand())
	fmt.Println("Model:", mercedes.Model())
	fmt.Println("Max Speed:", mercedes.MaxSpeed())
	fmt.Println("Engine Power:", mercedes.EnginePower())
	fmt.Println("Length in CM:", mercedes.Dimensions().Length().Get(CM))
	fmt.Println("Width in CM:", mercedes.Dimensions().Width().Get(CM))
	fmt.Println("Height in CM:", mercedes.Dimensions().Height().Get(CM))
	fmt.Println(strings.Repeat("-", 24))

	// BMW
	bmw := NewAutoBMW("Z4", 250, 252, 432.4, 186.7, 131.3)
	fmt.Println("Brand:", bmw.Brand())
	fmt.Println("Model:", bmw.Model())
	fmt.Println("Max Speed:", bmw.MaxSpeed())
	fmt.Println("Engine Power:", bmw.EnginePower())
	fmt.Println("Length in CM:", bmw.Dimensions().Length().Get(CM))
	fmt.Println("Width in CM:", bmw.Dimensions().Width().Get(CM))
	fmt.Println("Height in CM:", bmw.Dimensions().Height().Get(CM))
	fmt.Println(strings.Repeat("-", 24))

	// Dodge
	// dodge := NewAutoDodge("Viper", 275, 645, 176.6, 76.4, 49.0, Inch)
	dodge := NewAutoDodge("Viper", 275, 645, 448.5, 194.1, 124.5, CM)
	fmt.Println("Brand:", dodge.Brand())
	fmt.Println("Model:", dodge.Model())
	fmt.Println("Max Speed:", dodge.MaxSpeed())
	fmt.Println("Engine Power:", dodge.EnginePower())
	fmt.Println("Length in Inch:", dodge.Dimensions().Length().Get(Inch))
	fmt.Println("Width in Inch:", dodge.Dimensions().Width().Get(Inch))
	fmt.Println("Height in Inch:", dodge.Dimensions().Height().Get(Inch))
}
