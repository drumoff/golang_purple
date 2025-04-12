package main

import (
	"fmt"
	"math"
)

const IMTPower float64 = 2

func main() {
	

	fmt.Print(`
___ Калькулятор индекса массы тела ___
--------------------------------------	
`) // Мультистроки
    usrHeight, usrKg := getUserInput()
	IMT := calculateIMT(usrKg, usrHeight)
	isLean := IMT < 16
	if isLean {
		fmt.Println("У вас недостаток веса!")
	}
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userKg, userHeight float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (userHeight float64, userKg float64) {
	fmt.Println("Введите свой рост в сантиметрах")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес в киллограммах")
	fmt.Scan(&userKg)
	return
}