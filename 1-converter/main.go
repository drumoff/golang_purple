package main

import "fmt"

const USDtoEUR float64 = 1
const USDtoRUB float64 = 3
const EURtoRUB float64 = USDtoRUB / USDtoEUR

func main() {
	returnInputUser()
}

func returnInputUser() string {
	var input string
	fmt.Scan(&input)
	return input
}

func calcMoney(summ int, firstСurrency string, secCurrency string) float64 {
	var newSumm float64
	fmt.Println(summ)
	return newSumm
}