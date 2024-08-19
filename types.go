package azrp

// VmPrice is a type that reflects the price of an Azure VM within a specific
// region. It includes the VmSku, location, currently along with pricing for
// hourly PAYG, monthly 1 and 3 year reserved instance and 1 and 3 year saving
// plan.
type VmPrice struct {
	VmSku      string
	Region     string
	Currency   string
	PaygHrRate float32 // hourly payg rate
	OneYrRi    float32 // monthly rate for 1 year reserved instance
	ThreeYrRi  float32 // monthly rate for 3 year reserved instance
	OneYrSp    float32 // monthly rate for 1 year saving plan
	ThreeYrSp  float32 // monthly rate for 3 year saving plan
}

// PssdPrice is a type that represents the price of an Azure Premium Disk (v1)
// Currently, only LRS storage is supported
type PssdPrice struct {
	PssdName string  // name of the disk eg P10
	SizeGiB  uint    // size of disk in GiB
	Region   string  // location of the disk
	Currency string  // currency of pricing
	Price    float32 // monthly cost of disk
}

// Pssdv2Price is a struct that represent the pricing of an Azure Premium Disk
// v2. Pricing for a v2 disk is made up of three components, size, IOPS and
// throughput.
// The PriceGiB represents the cost for a 1 GiB of storage
// PriceIops represents the cost of 1 IOPS
// PriceMbs represents the cost of 1MiB/s throughput
// The idea is that the user can calculate the price of any pssdv2 disk using
// the returned data. However, the user should be aware the pssdv2 disk provide
// 3000 IOPS and and 125MB/s throughput as standard. Only additional IOPS and
// throughput is required to be priced.
// For a full list of limitations and constraints see -
// https://learn.microsoft.com/en-us/azure/virtual-machines/disks-types#premium-ssd-v2
type Pssdv2Price struct {
	Region    string  // region of the disk
	Currency  string  // currency of the pricing
	PriceGiB  float32 // price for 1GiB
	PriceIops float32 // price for 1 IOPS
	PriceMBs  float32 // price per MB/s
}
