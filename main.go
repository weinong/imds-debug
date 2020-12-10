package main

import (
	"flag"
	"fmt"

	"github.com/Azure/go-autorest/autorest/adal"
)

var (
	goVersion     string
	msiResourceID string
	resource      string
)

func init() {
	flag.StringVar(&msiResourceID, "msiResourceID", "", "msi resource ID")
	flag.StringVar(&resource, "resource", "https://management.azure.com/", "resource name")
}

func main() {
	flag.Parse()
	msiEndpoint, err := adal.GetMSIEndpoint()
	if err != nil {
		panic(err)
	}

	fmt.Println("go version: ", goVersion)
	fmt.Println("msiResourceID: ", msiResourceID)
	fmt.Println("resource: ", resource)
	fmt.Println("msiEndpoint: ", msiEndpoint)
	var spt *adal.ServicePrincipalToken
	if msiResourceID == "" {
		spt, err = adal.NewServicePrincipalTokenFromMSI(msiEndpoint, resource)
	} else {
		spt, err = adal.NewServicePrincipalTokenFromMSIWithIdentityResourceID(msiEndpoint, resource, msiResourceID)
	}
	if err != nil {
		panic(err)
	}

	spt.EnsureFresh()

	if spt.OAuthToken() != "" {
		fmt.Println("succeeded")
	}
}
