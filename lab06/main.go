package main

import (
	"fmt"
	"stockmarket/lab06/functions"
	"stockmarket/lab06/structures"
	"stockmarket/lab06/ui"
)

func loadAndProcessData(filename string) ([]structures.Data, error) {
	data, err := functions.LoadDataFromCSV(filename)
	if err != nil {
		return nil, fmt.Errorf("error loading data from %s: %w", filename, err)
	}
	return data, nil
}

func processDataForCalculation(calculationType string) {
	dataChoice := ui.ChooseData()
	var filename string
	switch dataChoice {
	case 1:
		filename = "./source/aapl.csv"
	case 2:
		filename = "./source/nflx.csv"
	case 3:
		filename = "./source/sbux.csv"
	default:
		fmt.Println("Invalid choice. Please try again.")
		return
	}

	data, err := loadAndProcessData(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	beginning, end := ui.ChooseDataPeriod()
	data = functions.ChoosePeriod(data, beginning, end)

	fmt.Println("Data for", calculationType)
	ui.ShowData(data)

	var result float64
	switch calculationType {
	case "RSI":
		result = functions.RSI(data)
	case "EMA":
		result = functions.EMA(data)
	case "ATR":
		result = functions.ATR(data)
	}

	fmt.Printf("%s: %.2f\n", calculationType, result)
}

func main() {
	for {
		choice := ui.ShowMainMenu()
		switch choice {
		case 1:
			dataChoice := ui.ChooseData()
			var filename string
			switch dataChoice {
			case 1:
				filename = "./source/aapl.csv"
			case 2:
				filename = "./source/nflx.csv"
			case 3:
				filename = "./source/sbux.csv"
			default:
				fmt.Println("Invalid choice. Please try again.")
				continue
			}

			data, err := loadAndProcessData(filename)
			if err != nil {
				fmt.Println(err)
				continue
			}
			ui.ShowData(data)

		case 2:
			method := ui.CalculateChoice()
			switch method {
			case 1:
				processDataForCalculation("RSI")
			case 2:
				processDataForCalculation("EMA")
			case 3:
				processDataForCalculation("ATR")
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
