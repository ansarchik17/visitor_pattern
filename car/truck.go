package car

type Truck struct {
	Model string
	Days  int
}

func (truck *Truck) GetType() string {
	return "Truck"
}

func (truck *Truck) Accept(visitor Visitor) {
	visitor.VisitTruck(truck)
}
