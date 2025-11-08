package car

type Car interface {
	GetType() string
	Accept(Visitor)
}

type Visitor interface {
	VisitSedan(*Sedan)
	VisitTruck(*Truck)
}
