package azrp

import (
	"fmt"
	"log/slog"
	"slices"
)

func (p Pricer) GetSssdPrice(name, region, currency string) (SssdPrice, error) {
	/* Ensure location is valid */
	if !slices.Contains(supReg, region) {
		return SssdPrice{}, fmt.Errorf("\"%s\" is not a valid azure region", region)
	}
	//Ensure pdisk name is valid
	if !slices.Contains(sdisks, name) {
		return SssdPrice{}, fmt.Errorf("\"%s\" is not a valid sdisk", name)
	}

	diskUrl, OpsUrl := getSssdStrings(name, region, currency)

	arDisk, err := p.apg(diskUrl)
	if err != nil {
		slog.Info("Failed to price sdisk")
		return SssdPrice{}, fmt.Errorf("failed to price sssd")
	}

	if len(arDisk.Items) != 1 || arDisk.Items == nil {
		if len(arDisk.Items) == 0 || arDisk.Items == nil {
			slog.Error("sssd API call got no results")
			return SssdPrice{}, fmt.Errorf("no results")
		}
		if len(arDisk.Items) > 1 {
			slog.Error("sssd API call got more than 1 disk result")
			return SssdPrice{}, fmt.Errorf("no results")
		}
	}

	arOps, err := p.apg(OpsUrl)
	if err != nil {
		slog.Info("Failed to price standard ssd disk")
		return SssdPrice{}, fmt.Errorf("failed to price sssd")
	}

	if len(arOps.Items) != 1 || arOps.Items == nil {
		if len(arOps.Items) == 0 || arOps.Items == nil {
			slog.Error("sssd ops API call got no results")
			return SssdPrice{}, fmt.Errorf("no results")
		}
		if len(arOps.Items) > 1 {
			slog.Error("sssd API call got more than 1 ops result")
			return SssdPrice{}, fmt.Errorf("no results")
		}
	}

	pp := SssdPrice{}

	pp.SssdName = name
	pp.Currency = currency
	pp.Region = region
	pp.Price = arDisk.Items[0].RetailPrice
	pp.OpsPrice = arOps.Items[0].RetailPrice
	size, _ := getSizeFromSssd(name)
	/* No need to check for error, the code wouldn't get this far is name was */
	/* incorrect                                                              */
	pp.SizeGiB = size

	return pp, nil
}
