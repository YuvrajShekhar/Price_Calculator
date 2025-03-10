package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrice        []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrice {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxrate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxrate,
	}
}
