package azrp

import (
	"fmt"
	"slices"
)

/* Preview version of the API currently needed for pssd_v2 prices*/

// URL for the Azure Retail Price API
const ApiUrl = "https://prices.azure.com/api/retail/prices"

// Preview Version used by azrp
const ApiPreview = "2023-01-01-preview"

// sapCur is a slice of strings the represent supported currencies
var supCur = []string{"USD",
	"AUD",
	"BRL",
	"CAD",
	"CHF",
	"CNY",
	"DKK",
	"EUR",
	"GBP",
	"INR",
	"JPY",
	"KRW",
	"NOK",
	"NZD",
	"RUB",
	"SEK",
	"TWD"}

/* list of regions generated with az */
/* az account list-locations --query "[].{Name:name}" -o table */

// supReg is a slice of strings that represent the supported Azure regions
var supReg = []string{
	"eastus",
	"southcentralus",
	"westus2",
	"westus3",
	"australiaeast",
	"southeastasia",
	"northeurope",
	"swedencentral",
	"uksouth",
	"westeurope",
	"centralus",
	"southafricanorth",
	"centralindia",
	"eastasia",
	"japaneast",
	"koreacentral",
	"canadacentral",
	"francecentral",
	"germanywestcentral",
	"italynorth",
	"norwayeast",
	"polandcentral",
	"spaincentral",
	"switzerlandnorth",
	"mexicocentral",
	"uaenorth",
	"brazilsouth",
	"israelcentral",
	"qatarcentral",
	"centralusstage",
	"eastusstage",
	"eastus2stage",
	"northcentralusstage",
	"southcentralusstage",
	"westusstage",
	"westus2stage",
	"asia",
	"asiapacific",
	"australia",
	"brazil",
	"canada",
	"europe",
	"france",
	"germany",
	"global",
	"india",
	"israel",
	"italy",
	"japan",
	"korea",
	"newzealand",
	"norway",
	"poland",
	"qatar",
	"singapore",
	"southafrica",
	"sweden",
	"switzerland",
	"uae",
	"uk",
	"unitedstates",
	"unitedstateseuap",
	"eastasiastage",
	"southeastasiastage",
	"brazilus",
	"eastus2",
	"eastusstg",
	"northcentralus",
	"westus",
	"japanwest",
	"jioindiawest",
	"centraluseuap",
	"eastus2euap",
	"westcentralus",
	"southafricawest",
	"australiacentral",
	"australiacentral2",
	"australiasoutheast",
	"jioindiacentral",
	"koreasouth",
	"southindia",
	"westindia",
	"canadaeast",
	"francesouth",
	"germanynorth",
	"norwaywest",
	"switzerlandwest",
	"ukwest",
	"uaecentral",
	"brazilsoutheast"}

var pdisks []string = []string{
	"P1",
	"P2",
	"P3",
	"P4",
	"P6",
	"P10",
	"P15",
	"P20",
	"P30",
	"P40",
	"P50",
	"P60",
	"P70",
	"P80",
}

// GetPssdFromSize is a function that takes the size of a required premium SSD
// disk and returns the SKU string for pricing. This allows users who do not
// know SKU of the disk sizes to simply supply the disk size in the config.
// It can also be used inline with GetPssdPrice for example
// `diskPrice, err := p.GetPssdPrice(azrp.GetPssdFromSize(128), "uksouth")`
func GetPssdFromSize(sz uint) string {
	switch {
	case sz <= 4:
		return "P1"
	case sz <= 8:
		return "P2"
	case sz <= 16:
		return "P3"
	case sz <= 32:
		return "P4"
	case sz <= 64:
		return "P6"
	case sz <= 128:
		return "P10"
	case sz <= 256:
		return "P15"
	case sz <= 512:
		return "P20"
	case sz <= 1024:
		return "P30"
	case sz <= 2048:
		return "P40"
	case sz <= 4096:
		return "P50"
	case sz <= 8192:
		return "P60"
	case sz <= 16384:
		return "P70"
	case sz <= 32768:
		return "P80"
	default:
		return "error"
	}
}

func getSizeFromPssd(pssd string) (uint, error) {
	switch pssd {
	case "P1":
		return 4, nil
	case "P2":
		return 8, nil
	case "P3":
		return 16, nil
	case "P4":
		return 32, nil
	case "P6":
		return 64, nil
	case "P10":
		return 128, nil
	case "P15":
		return 256, nil
	case "P20":
		return 512, nil
	case "P30":
		return 1024, nil
	case "P40":
		return 2048, nil
	case "P50":
		return 4096, nil
	case "P60":
		return 8192, nil
	case "P70":
		return 16384, nil
	case "P80":
		return 32768, nil
	default:
		return 0, fmt.Errorf("not a valid pdisk name")
	}
}

func validateCurrencyCode(cur string) bool {
	return slices.Contains(supCur, cur)
}

func validateLocation(loc string) bool {
	return slices.Contains(supReg, loc)
}
