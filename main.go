package main

import (
	"assik8_Visitor/car"
	"assik8_Visitor/services"
	"fmt"
)

func main() {
	sedanHonda := &car.Sedan{Model: "Honda", Days: 3}
	truckFord := &car.Truck{Model: "Ford", Days: 2}

	rentCalculator := &services.RentCalculator{}
	maintenanceChecker := &services.MaintenanceChecker{}

	fmt.Println("1)Rental Cost Calculation!")
	sedanHonda.Accept(rentCalculator)
	truckFord.Accept(rentCalculator)

	fmt.Println("\n2)Maintenance Check!")
	sedanHonda.Accept(maintenanceChecker)
	truckFord.Accept(maintenanceChecker)
}
