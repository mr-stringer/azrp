package azrp

import (
	"fmt"
	"log/slog"
	"slices"
)

func (p Pricer) GetPssdPrice(name, region, currency string) (PssdPrice, error) {
	/* Ensure location is valid */
	if !slices.Contains(supReg, region) {
		return PssdPrice{}, fmt.Errorf("\"%s\" is not a valid azure region", region)
	}
	//Ensure pdisk name is valid
	if !slices.Contains(pdisks, name) {
		return PssdPrice{}, fmt.Errorf("\"%s\" is not a valid pdisk", name)
	}
	p.ApiVersion = ApiPreview
	p.Currency = currency
	p.ArmRegionName = region
	p.MeterName = fmt.Sprintf("%s LRS Disk", name)
	p.ProductName = "Premium SSD Managed Disks"
	p.SkuName = fmt.Sprintf("%s LRS", name)
	p.ServiceFamily = "Storage"
	p.PriceType = "Consumption"

	s1, err := p.GetString()

	if err != nil {
		return PssdPrice{}, err
	}

	ar, err := p.apg(s1)
	if err != nil {
		slog.Info("Failed to price pdisk")
		return PssdPrice{}, fmt.Errorf("failed to price pssd")
	}

	if len(ar.Items) != 1 || ar.Items == nil {
		if len(ar.Items) == 0 || ar.Items == nil {
			slog.Error("pssd API call got no results")
			return PssdPrice{}, fmt.Errorf("no results")
		}
		if len(ar.Items) > 1 {
			slog.Error("pssd API call got more than 1 result")
			return PssdPrice{}, fmt.Errorf("no results")
		}
	}

	pp := PssdPrice{}

	pp.PssdName = name
	pp.Currency = currency
	pp.Region = region
	pp.Price = ar.Items[0].RetailPrice
	/*No need to check for error, disk name cannot be wrong here */
	size, _ := GetSizeFromPssd(name)
	pp.SizeGiB = size

	return pp, nil
}
