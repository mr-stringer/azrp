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

	fmt.Print(getPssdString(name, region, currency))
	ar, err := p.apg(getPssdString(name, region, currency))
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
	size, err := getSizeFromPssd(name)
	if err != nil {
		/* These lines can't be easily tested as we already check that incorrect names */
		/* at the top of the function */
		slog.Error("Pdisk name cannot be resolved to a size", "pdisk name", name)
		return PssdPrice{}, fmt.Errorf("pdisk name %s cannot be resolved to a size", name)
	}
	pp.SizeGiB = size

	return pp, nil
}
