package car

type Sedan struct {
	Model string
	Days  int
}

func (sedan *Sedan) GetType() string {
	return sedan.Model
}

func (sedan *Sedan) Accept(visitor Visitor) {
	visitor.VisitSedan(sedan)
}
