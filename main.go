package main

import "fmt"

func main() {

	var prices []float64 = []float64{10, 1, 14.5, 89}
	taxRate := []float64{1.3, 4.5, 0.8, 8.5, 4.8}

	//var result map[float64] []float64 = make(map[float64] []float64)
	result := make(map[float64][]float64)

	for _, taxValue := range taxRate {

		pricesAfterTax := make([]float64, len(prices))
		for priceIndex, priceValue := range prices {
			pricesAfterTax[priceIndex] = priceValue * taxValue
		}
		result[taxValue] = pricesAfterTax
	}
	fmt.Println(result)

}
