package azrp

import (
	"fmt"
	"log/slog"
	"strings"
)

func getPssdString(pdiskName, region, currency string) string {
	s1 := fmt.Sprintf("%s?api-version=%s&CurrencyCode=%s&$filter=armRegionName eq '%s' and serviceFamily eq 'Storage' and skuName eq '%s LRS' and productName eq 'Premium SSD Managed Disks' and meterName eq '%s LRS Disk' and priceType eq 'Consumption'", ApiUrl, ApiPreview, currency, region, pdiskName, pdiskName)
	/* Need url encoding */
	/* Due to MS API weirdness, can't use net/url */
	/* but the following does what is needed */
	s1 = strings.Replace(s1, " ", "%20", -1)
	s1 = strings.Replace(s1, "'", "%27", -1)
	slog.Debug("ApiPssdPriceString", "url", s1)
	return s1
}

func getPssdv2String(region, currency string) string {
	s1 := fmt.Sprintf("%s?api-version=%s&CurrencyCode=%s&$filter=armRegionName eq '%s' and serviceFamily eq 'Storage' and priceType eq 'Consumption' and productName eq 'Azure Premium SSD v2'", ApiUrl, ApiPreview, currency, region)
	/* Need url encoding */
	/* Due to MS API weirdness, can't use net/url */
	/* but the following does what is needed */
	s1 = strings.Replace(s1, " ", "%20", -1)
	slog.Debug("ApiPssdv2PriceString", "url", s1)
	return s1
}

func getVmString(vm, loc, cur string) string {
	s1 := fmt.Sprintf("%s?api-version=%s&CurrencyCode=%s&$filter=armRegionName eq '%s' and serviceFamily eq 'Compute' and serviceName eq 'Virtual Machines' and armSkuName eq '%s'", ApiUrl, ApiPreview, cur, loc, vm)
	/* Need url encoding */
	/* Due to MS API weirdness, can't use net/url */
	/* but the following does what is needed */
	s1 = strings.Replace(s1, " ", "%20", -1)
	s1 = strings.Replace(s1, "'", "%27", -1)
	return s1
}
