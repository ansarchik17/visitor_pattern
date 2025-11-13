package car

type Truck struct {
	Model string
	Days  int
}

func (truck *Truck) GetType() string {
	return truck.Model
}

func (truck *Truck) Accept(visitor Visitor) {
	visitor.VisitTruck(truck)
}
