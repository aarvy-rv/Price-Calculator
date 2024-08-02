package main

import (
	"fmt"
	"strings"

	"practice.com/price-calculator/filemanager"
	"practice.com/price-calculator/price"
)

func main() {

	taxRate := []float64{1.3, 4.5, 0.8, 8.5, 4.8}

	//var result map[float64] []float64 = make(map[float64] []float64)

	for _, taxValue := range taxRate {
		var outputFileName string = "result" + strings.ReplaceAll(fmt.Sprintf("%f", taxValue), ".", "_") + ".json"
		fm := filemanager.NewFileManger("prices.txt", outputFileName)
		job := price.NewTaxIncludedPriceJob(*fm, taxValue)
		job.PriceProcessing()
	}

}
