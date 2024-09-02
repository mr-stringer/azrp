package azrp

import "fmt"

func (p Pricer) Execute() (*ApiResponse, error) {
	if p.ArmRegionName == "" {
		if !validateLocation(p.ArmRegionName) {
			return &ApiResponse{}, fmt.Errorf("unsupported ArmRegionName")
		}
	}
	if p.Currency == "" {
		return &ApiResponse{}, fmt.Errorf("currency not set")
	}
	if !validateCurrencyCode(p.Currency) {
		return &ApiResponse{}, fmt.Errorf("unsupported currency")
	}

	s1, err := p.GetString()
	if err != nil {
		return &ApiResponse{}, err
	}

	ar, err := p.apg(s1)
	if err != nil {
		return &ApiResponse{}, err
	}

	return &ar, nil
}
