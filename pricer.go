package azrp

type Pricer struct {
	apg           apiGetter
	Currency      string
	ApiVersion    string
	ArmRegionName string
	Location      string
	MeterId       string
	MeterName     string
	ProductId     string
	SkuId         string
	ProductName   string
	SkuName       string
	ServiceName   string
	ServiceId     string
	ServiceFamily string
	PriceType     string
	ArmSkuName    string
}

func NewPricer() *Pricer {
	return &Pricer{apg: apiGet}
}
