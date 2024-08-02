package price

import (
	"fmt"
	"math"

	"practice.com/price-calculator/conversion"
	"practice.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"` //excludes this field in json
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.Read()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) PriceProcessing() error {

	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = math.Round((price*(1+job.TaxRate))*100) / 100 //Rounding to two decimal places
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResults(job)
}

func NewTaxIncludedPriceJob(ioManager iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   ioManager,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 1, 14.5, 89},
	}
}
