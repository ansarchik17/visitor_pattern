package services

import (
	"assik8_Visitor/car"
	"fmt"
)

type MaintenanceChecker struct{}

func (maintenance *MaintenanceChecker) VisitSedan(sedan *car.Sedan) {
	fmt.Printf("Maintenance check for Sedan %s: Oil change required.\n", sedan.Model)
}

func (maintenance *MaintenanceChecker) VisitTruck(truck *car.Truck) {
	fmt.Printf("Maintenance check for Truck %s: Brake inspection due.\n", truck.Model)
}
