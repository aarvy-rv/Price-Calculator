package cmdmanager

import "fmt"

type CMDManager struct {
}

func NewCmdManager() CMDManager {
	return CMDManager{}
}

func (cmd CMDManager) Read() ([]string, error) {

	fmt.Println("Please Enter your prices . Confirm each price with enter")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}
func (cmd CMDManager) WriteResults(data interface{}) error {
	fmt.Println(data)
	return nil

}
