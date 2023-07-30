package main

import (
	"fmt"

	arvan "github.com/arvansdk/go/arvan"
)

func main() {
	client := arvan.NewClient()
	report := client.GetReports("linkbe.site")
	fmt.Println(report.Data.Statistics.Traffics.Total)
}