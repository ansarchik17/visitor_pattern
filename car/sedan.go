package car

type Sedan struct {
	Model string
	Days  int
}

func (sedan *Sedan) GetType() string {
	return "Sedan"
}

func (sedan *Sedan) Accept(v Visitor) {
	v.VisitSedan(sedan)
}
