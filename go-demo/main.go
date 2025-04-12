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

	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела!")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела!")
	case IMT < 25:
		fmt.Println("У вас нормальный вес!")
	case IMT < 30:
		fmt.Println("У вас избыточный вес!")
	default:
		fmt.Println("Ты жирдяй!")
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
