package azrp

import "fmt"

// GetString returns a string that is the url for the configuration of the
// ApiRequest type.
func (p Pricer) GetString() (string, error) {
	/* Check the basics */
	/* Check that currency, is correct and check if a preview version */
	/* needed */
	if p.Currency == "" {
		return "", fmt.Errorf("currency not set")
	}
	if !validateCurrencyCode(p.Currency) {
		return "", fmt.Errorf("currency code '%s' not supported", p.Currency)
	}

	var s1 string = ApiUrl

	if p.ApiVersion != "" {
		s1 = fmt.Sprintf("%s?api-version=%s&CurrencyCode=%s", s1, p.ApiVersion, p.Currency)
	} else {
		s1 = fmt.Sprintf("%s?CurrencyCode=%s", s1, p.Currency)
	}

	type filters struct {
		Key   string
		Value string
	}
	f1 := []filters{}

	if p.ArmRegionName != "" {
		f1 = append(f1, filters{"armRegionName", p.ArmRegionName})
	}
	if p.Location != "" {
		f1 = append(f1, filters{"location", p.Location})
	}
	if p.MeterId != "" {
		f1 = append(f1, filters{"meterId", p.MeterId})
	}
	if p.MeterName != "" {
		f1 = append(f1, filters{"meterName", p.MeterName})
	}
	if p.ProductId != "" {

		f1 = append(f1, filters{"productId", p.ProductId})
	}
	if p.SkuId != "" {
		f1 = append(f1, filters{"skuId", p.SkuId})
	}
	if p.ProductName != "" {
		f1 = append(f1, filters{"productName", p.ProductName})
	}
	if p.SkuName != "" {
		f1 = append(f1, filters{"skuName", p.SkuName})
	}
	if p.ServiceName != "" {
		f1 = append(f1, filters{"serviceName", p.ServiceName})
	}
	if p.ServiceId != "" {
		f1 = append(f1, filters{"serviceId", p.ServiceId})
	}
	if p.ServiceFamily != "" {
		f1 = append(f1, filters{"serviceFamily", p.ServiceFamily})
	}
	if p.PriceType != "" {
		f1 = append(f1, filters{"priceType", p.PriceType})
	}
	if p.ArmSkuName != "" {
		f1 = append(f1, filters{"armSkuName", p.ArmSkuName})
	}

	if len(f1) == 0 {
		return s1, nil
	}

	/* When more than one filter is present, they will always appear in the   */
	/* order in which that are added, this means the testing add multiple     */
	/* is always stable.                                                      */

	var first = true
	for _, v := range f1 {
		if first {
			s1 = fmt.Sprintf("%s&$filter=%s eq '%s'", s1, v.Key, v.Value)
			first = false
		} else {
			s1 = fmt.Sprintf("%s and %s eq '%s'", s1, v.Key, v.Value)
		}
	}

	return basicEncoder(s1), nil
}
