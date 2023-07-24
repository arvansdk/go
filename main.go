package main

import (
	arvan "github.com/arvansdk/go/arvan"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cdnClient := &arvan.APIClient{Service: "CDN"}
	testDomain := "jikopik.site"
	// info := cdnClient.GetDomainInfo("linkbe.site")
	cdnClient.CreateDomain(testDomain)
	cdnClient.AddDNSRecord(testDomain, "test", "2.2.2.2")
	cdnClient.UpdateSSLConfig(testDomain)
	// fmt.Println(info)
}