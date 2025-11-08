package visitor

import (
	"assik8_Visitor/car"
	"fmt"
)

type RentCalculator struct{}

func (rent *RentCalculator) VisitSedan(sedan *car.Sedan) {
	fmt.Printf("Rental for Sedan %s: $%d\n", sedan.Model, sedan.Days*40)
}

func (rent *RentCalculator) VisitTruck(truck *car.Truck) {
	fmt.Printf("Rental for Truck %s: $%d\n", truck.Model, truck.Days*80)
}
