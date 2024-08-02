package price

import (
	"fmt"
	"math"
	"strings"

	"practice.com/price-calculator/conversion"
	"practice.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadFile("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) PriceProcessing() {

	job.LoadData()

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = math.Round((price*(1+job.TaxRate))*100) / 100 //Rounding to two decimal places
	}

	job.TaxIncludedPrices = result

	var outputFileName string = "result" + strings.ReplaceAll(fmt.Sprintf("%f", job.TaxRate), ".", "_") + ".json"
	err := filemanager.WriteToJson(job, outputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 1, 14.5, 89},
	}
}
