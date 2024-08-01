package main

import "practice.com/price-calculator/price"

func main() {

	taxRate := []float64{1.3, 4.5, 0.8, 8.5, 4.8}

	//var result map[float64] []float64 = make(map[float64] []float64)

	for _, taxValue := range taxRate {
		job := price.NewTaxIncludedPriceJob(taxValue)
		job.PriceProcessing()
	}

}
