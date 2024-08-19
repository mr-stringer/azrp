package azrp

type Pricer struct {
	apg apiGetter
}

func NewPricer() *Pricer {
	return &Pricer{apg: apiGet}
}
