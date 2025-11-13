package services

import (
	"assik8_Visitor/car"
	"fmt"
)

type RentCalculator struct{}

func (rent *RentCalculator) VisitSedan(sedan *car.Sedan) {
	totalSedanCost := sedan.Days * 40
	fmt.Printf("Rental for Sedan %s: $%d\n", sedan.Model, totalSedanCost)
}

func (rent *RentCalculator) VisitTruck(truck *car.Truck) {
	totalTruckCost := truck.Days * 80
	fmt.Printf("Rental for Truck %s: $%d\n", truck.Model, totalTruckCost)
}
