package main

import (
	arvan "github.com/arvansdk/go/arvan"
)

func main() {
	cdnClient := &arvan.APIClient{}
	testDomain := "jikopik.site"
	// info := cdnClient.GetDomainInfo("linkbe.site")
	cdnClient.CreateDomain(testDomain)
	cdnClient.AddDNSRecord(testDomain, "test", "2.2.2.2")
	cdnClient.UpdateSSLConfig(testDomain)
	// fmt.Println(info)
}