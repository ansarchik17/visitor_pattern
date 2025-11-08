package main

import (
	"assik8_Visitor/car"
	"assik8_Visitor/visitor"
	"fmt"
)

func main() {
	sedanHonda := &car.Sedan{Model: "Honda", Days: 3}
	truckFord := &car.Truck{Model: "Ford", Days: 2}

	rentCalc := &visitor.RentCalculator{}
	maintCheck := &visitor.MaintenanceChecker{}

	fmt.Println("1)Rental Cost Calculation!")
	sedanHonda.Accept(rentCalc)
	truckFord.Accept(rentCalc)

	fmt.Println("\n2)Maintenance Check!")
	sedanHonda.Accept(maintCheck)
	truckFord.Accept(maintCheck)
}
