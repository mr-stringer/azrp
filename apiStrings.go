package azrp

import (
	"fmt"
	"log/slog"
	"strings"
)

// getSssdStrings returns url strings that are used to price Standard SSD disks
// The first is used to get the url for pricing the disk, the second is used
// to get the pricing for disk operations
func getSssdStrings(sssdName, region, currency string) (disk, ops string) {
	disk = fmt.Sprintf("%s?api-version=%s&CurrencyCode=%s&$filter=armRegionName eq '%s' and serviceFamily eq 'Storage' and productName eq 'Standard SSD Managed Disks' and skuName eq '%s LRS' and meterName eq '%s LRS Disk' and priceType eq 'Consumption'", ApiUrl, ApiPreview, currency, region, sssdName, sssdName)
	ops = fmt.Sprintf("%s?api-version=%s&CurrencyCode=%s&$filter=armRegionName eq '%s' and serviceFamily eq 'Storage' and productName eq 'Standard SSD Managed Disks' and skuName eq '%s LRS' and meterName eq '%s LRS Disk Operations' and priceType eq 'Consumption'", ApiUrl, ApiPreview, currency, region, sssdName, sssdName)
	disk = basicEncoder(disk)
	ops = basicEncoder(ops)
	slog.Debug(disk)
	slog.Debug(ops)
	return
}

//func getPssdString(pdiskName, region, currency string) string {
//	s1 := fmt.Sprintf("%s?api-version=%s&CurrencyCode=%s&$filter=armRegionName eq '%s' and serviceFamily eq 'Storage' and skuName eq '%s LRS' and productName eq 'Premium SSD Managed Disks' and meterName eq '%s LRS Disk' and priceType eq 'Consumption'", ApiUrl, ApiPreview, currency, region, pdiskName, pdiskName)
//	/* Need url encoding */
//	/* Due to MS API weirdness, can't use net/url */
//	/* but the following does what is needed */
//	s1 = basicEncoder(s1)
//	slog.Debug("ApiPssdPriceString", "url", s1)
//	return s1
//}

//func getPssdv2String(region, currency string) string {
//s1 := fmt.Sprintf("%s?api-version=%s&CurrencyCode=%s&$filter=armRegionName eq '%s' and serviceFamily eq 'Storage' and priceType eq 'Consumption' and productName eq 'Azure Premium SSD v2'", ApiUrl, ApiPreview, currency, region)
//	/* Need url encoding */
//	/* Due to MS API weirdness, can't use net/url */
//	/* but the following does what is needed */
//	s1 = basicEncoder(s1)
//	slog.Debug("ApiPssdv2PriceString", "url", s1)
//	return s1
//}

//func getVmString(vm, loc, cur string) string {
//	s1 := fmt.Sprintf("%s?api-version=%s&CurrencyCode=%s&$filter=armRegionName eq '%s' and serviceFamily eq 'Compute' and serviceName eq 'Virtual Machines' and armSkuName eq '%s'", ApiUrl, ApiPreview, cur, loc, vm)
//	/* Need url encoding */
//	/* Due to MS API weirdness, can't use net/url */
//	/* but the following does what is needed */
//	s1 = basicEncoder(s1)
//	return s1
//}

func basicEncoder(s1 string) string {
	s1 = strings.Replace(s1, " ", "%20", -1)
	s1 = strings.Replace(s1, "'", "%27", -1)
	return s1
}
