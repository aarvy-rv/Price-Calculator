package price

import (
	"fmt"
	"math"

	"practice.com/price-calculator/conversion"
	"practice.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadFile()
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

	err := job.IOManager.WriteToJson(job)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewTaxIncludedPriceJob(ioManager filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   ioManager,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 1, 14.5, 89},
	}
}
