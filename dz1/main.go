package main

import (
	"errors"
	"fmt"
)

const USDtoRUB float64 = 82.5
const USDtoEUR float64 = 0.88
const EURtoRUB float64 = USDtoRUB / USDtoEUR



func main() {
	fmt.Println("===============\nКонвертер валют\n===============")

	for i := 0; i < 1; i++ {
		frstCurrency, sum, scndCurrency := getUserInput()

	}
}


func getUserInput() (frstCurrency string, sum float64, scndCurrency string) {
	fmt.Print("Введите исходную валюту(RUB/USD/EUR): ")
    fmt.Scan(&frstCurrency)
	fmt.Print("Введите сумму: ")
    fmt.Scan(&sum)
	fmt.Print("Введите целевую валюту(RUB/USD/EUR): ")
	fmt.Scan(&scndCurrency)
	return
}

func calculateCurency(frstCurrency string, scndCurrency string, sum float64) (float64, error) {
	fmt.Println(frstCurrency)
	inputSum := fmt.Sprintf("%.0f", sum)
	fmt.Println(inputSum)
	fmt.Println(scndCurrency)
     
	
	return result, nil
}