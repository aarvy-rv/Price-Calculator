package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(input []string) ([]float64, error) {

	var output = make([]float64, len(input))

	for lineIndex, line := range input {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("error in converting string to float64")
			return nil, errors.New("failed to convert string to float")

		}
		output[lineIndex] = price
	}
	return output, nil

}
