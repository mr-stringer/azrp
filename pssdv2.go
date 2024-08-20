package azrp

import (
	"fmt"
	"log/slog"
)

// The function GetPssdv2Price takes a region and a currency and returns */
// a Pssdv2Price struct */
// Unlike pssds, this function returns PER HOUR pricing, not monthly costs!
func (p Pricer) GetPssdv2Price(region, currency string) (Pssdv2Price, error) {
	if !validateLocation(region) {
		return Pssdv2Price{}, fmt.Errorf("unsupported location")
	}
	if !validateCurrencyCode(currency) {
		return Pssdv2Price{}, fmt.Errorf("unsupported currency")
	}

	ar, err := p.apg(getPssdv2String(region, currency))
	if err != nil {
		slog.Error("Failed to price pssdv2")
		return Pssdv2Price{}, fmt.Errorf("failed to price pssdv2")
	}

	/* Check that items is populated */
	if len(ar.Items) == 0 || ar.Items == nil {
		slog.Error("pssdv2 API call got no results")
		return Pssdv2Price{}, fmt.Errorf("no results")
	}

	dp := Pssdv2Price{}
	dp.Currency = currency
	dp.Region = region

	/* Loop over result */
	for _, item := range ar.Items {
		/* immediately discard if price is empty */
		switch {
		case item.RetailPrice == 0:
			continue
		case item.MeterName == "Premium LRS Provisioned IOPS":
			dp.PriceIops = item.RetailPrice
		case item.MeterName == "Premium LRS Provisioned Throughput (MBps)":
			dp.PriceMBs = item.RetailPrice
		case item.MeterName == "Premium LRS Provisioned Capacity":
			dp.PriceGiB = item.RetailPrice
		}
	}

	/* Ensure that all parts of the disk were priced */
	if dp.PriceIops == 0 {
		slog.Error("Unable to get IOPS for pssdv2", "region", region, "currency", currency)
		return Pssdv2Price{}, fmt.Errorf("unable to get IOPS price pssdv2")
	}
	if dp.PriceMBs == 0 {
		slog.Error("Unable to get throughput price for pssdv2", "region", region, "currency", currency)
		return Pssdv2Price{}, fmt.Errorf("unable to get throughput price pssdv2")
	}
	if dp.PriceGiB == 0 {
		slog.Error("Unable to get capacity price for pssdv2", "region", region, "currency", currency)
		return Pssdv2Price{}, fmt.Errorf("unable to get capacity price pssdv2")
	}

	return dp, nil
}
