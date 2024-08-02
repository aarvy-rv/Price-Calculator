package price

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {

	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not read file!")
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
		fmt.Println("Error in Reading Line")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("error in converting string to float64")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIndex] = price
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) PriceProcessing() {

	job.LoadData()

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = math.Round((price*(1+job.TaxRate))*100) / 100 //Rounding to two decimal places
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 1, 14.5, 89},
	}
}
