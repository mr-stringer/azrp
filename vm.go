package azrp

import (
	"fmt"
	"log/slog"
	"strings"
)

// GetVmPrice provides the price of a VM in a specific location in a specific
// currency.
func (p Pricer) GetVmPrice(location, currency, vmsku string) (VmPrice, error) {
	// Check currency code
	if !validateCurrencyCode(currency) {
		return VmPrice{}, fmt.Errorf("unsupported currency")
	}
	if !validateLocation(location) {
		return VmPrice{}, fmt.Errorf("unsupported location")
	}

	ar, err := p.apg(getVmString(vmsku, location, currency))
	if err != nil {
		slog.Error("Failed to price VM")
		return VmPrice{}, fmt.Errorf("failed to price vm")
	}

	/* Ensure that prices were returned */
	if len(ar.Items) == 0 || ar.Items == nil {
		slog.Error("VM API call got no results")
		return VmPrice{}, fmt.Errorf("no results")
	}

	vmp := VmPrice{
		VmSku:    vmsku,
		Region:   location,
		Currency: currency,
	}

	for _, v := range ar.Items {
		/* Skip all Windows versions */
		if strings.Contains(v.ProductName, "Windows") {
			continue
		}
		/* Skip all low priority and spot systems*/
		if strings.Contains(v.MeterName, "Low Priority") {
			continue
		}
		if strings.Contains(v.MeterName, "Spot") {
			continue
		}
		/* skip Cloud Services products (VMs only) */
		if strings.Contains(v.ProductName, "Cloud Services") {
			continue
		}
		if v.Type == "Consumption" {
			vmp.PaygHrRate = v.UnitPrice
		} else if v.Type == "Reservation" && v.ReservationTerm == "1 Year" {
			/*The price given is for the 1 year term, divide by 12 to get monthly */
			vmp.OneYrRi = (v.UnitPrice / 12)
		} else if v.Type == "Reservation" && v.ReservationTerm == "3 Years" {
			/*The price given is for the 3 year term, divide by 36 to get monthly */
			vmp.ThreeYrRi = (v.UnitPrice / 36)
		}
	}

	/* Ensure all fields are populated */
	if vmp.OneYrRi == 0 {
		slog.Error("Couldn't retrieve 1 Year RI price", "VmSku", vmsku, "Region", location)
		return VmPrice{}, fmt.Errorf("could not retrieve 1 year RI price for VM")
	}
	if vmp.ThreeYrRi == 0 {
		slog.Error("Couldn't retrieve 3 Year RI price", "VmSku", vmsku, "Region", location)
		return VmPrice{}, fmt.Errorf("could not retrieve 3 year RI price for VM")
	}
	if vmp.PaygHrRate == 0 {
		slog.Error("Couldn't retrieve hourly PAYG price", "VmSku", vmsku, "Region", location)
		return VmPrice{}, fmt.Errorf("could not retrieve hourly PAYG price for VM")
	}

	return vmp, nil

}
