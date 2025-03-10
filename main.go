package main

import (
	"example.com/price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxrate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxrate)
		priceJob.Process()
	}

}
