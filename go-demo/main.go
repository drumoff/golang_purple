package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(`
___ Калькулятор индекса массы тела ___
--------------------------------------	
`) // Мультистроки
    usrHeight, usrKg := getUserInput()
	IMT := calculateIMT(usrKg, usrHeight)
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

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Println("Введите свой рост в сантиметрах")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес в киллограммах")
	fmt.Scan(&userKg)
	return userHeight, userKg
}