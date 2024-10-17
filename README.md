# azrp

![badge](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/mr-stringer/087cd593d18ffd439a1acbd2576ea388/raw/go-cover-badge.json)

azrp is a go module for retrieving Azure retails prices.

The module allows users to retrieve Azure retails prices without any prior 
knowledge of the Azure retail pricing API.

The module provides a client which supports a number of functions that allow
users to access Azure pricing.

The module can be used to provide users with a simplified interface to the 
Azure Retail Price API, allowing them to easily create custom API requests.
However, for common Azure objects, the user may take advantage of simple helpers
that provide pricing for specific object types.

## Installation

Installing azrp

```shell
go get -u github.com/mr-stringer/azrp
```

## Initialising the client

To use azrp in your project you will usually create a new client.

```go
azrpClient := azrp.NewPricer()
```

## Helpers for common products

There are currently four helpers for common Azure products, there are: 

* Virtual Machines
* Standard SSD storage
* Premium SSD storage
* Premium SSD Storage V2

The helpers are easy to use and each return a specific type. The below
example shows how a pricing information for an Azure Virtual Machines can be 
obtained and used.

```go
prc := azrp.NewPricer()
vmp, err := prc.GetVmPrice("Standard_E20ds_v5", "uksouth", "GBP")
if err != nil {
	log.Fatal(err)
}

fmt.Printf("VM:%s\n", vmp.VmSku)
fmt.Printf("Region:%s\n", vmp.Region)
fmt.Printf("Currency:%s\n", vmp.Currency)
fmt.Printf("HourlyPayg:%0.4f\n", vmp.PaygHrRate)
fmt.Printf("1YrRiPcm:%0.4f\n", vmp.OneYrRi)
fmt.Printf("3YrRiPcm:%0.4f\n", vmp.ThreeYrRi)
```

The other helper functions of the Pricer type are:

* GetSssdPrice, used for getting Standard SSD prices
* GetPssdPrice, used for getting Premium SSD v1 prices
* GetPssdv2Price, used for getting Premium SSD v2 prices

See the go-docs for full documentation.

## Custom API request

The Pricer client has two functions that can be used to make customer API
requests. These are `GetString` and `Execute`. Both of these functions require
the Pricer client fields to be populated with the custom requirement.

The Pricer type is shown here:

```go
type Pricer struct {
	apg           apiGetter // used for testing
	Currency      string    // mandatory! 
	ApiVersion    string    // specifies specific API version 
	ArmRegionName string 
	Location      string
	MeterId       string
	MeterName     string
	ProductId     string
	SkuId         string
	ProductName   string
	SkuName       string
	ServiceName   string
	ServiceId     string
	ServiceFamily string
	PriceType     string
	ArmSkuName    string
}
```

To configure the relevant fields of the Pricer, follow the below example, in 
which, we want to acquire the prices of all storage products in the region
'germanywestcentral' in Euros.

```go
prc := azrp.NewPricer()
prc.Currency = "EUR"
prc.ArmRegionName = "germanywestcentral"
prc.ServiceFamily = "Storage"
```

With the Pricer configured, the user may run `GetString` or `Execute`.

`GetString` will return the full API endpoint as a string that can then be
passed to another process (such as http.Get).

However, running `Execute` will make the request and return the type
`ApiResponse`.

Care should be taken when creating and using custom API requests. The `Execute`
function simply marshals the result from the API into the `ApiResonse` type.
The results of the API return may not match what the user expects.

Also be aware the that further calls may be required to retrieve all elements.
If the field `ApiResonse.NextPageLink` is not an empty string, then the value
will be a new endpoint that will provide additional elements and possible
another endpoint.
