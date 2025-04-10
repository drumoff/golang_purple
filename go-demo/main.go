package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64
	var userKg float64
	fmt.Print(`
___ Калькулятор индекса массы тела ___
--------------------------------------	
`) // Мультистроки
	fmt.Println("Введите свой рост в сантиметрах")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес в киллограммах")
	fmt.Scan(&userKg)
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userKg, userHeight float64) float64 {
	const IMTPower float64 = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}