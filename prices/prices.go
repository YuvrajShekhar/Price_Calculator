package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrice        []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reaading the file content failed")
		fmt.Println("err")
		return
	}
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
