package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2
	var userHeight float64
	var userKg float64
	fmt.Print(`___ Калькулятор индекса массы тела ___`) // Мультистроки
	fmt.Println("Введите свой рост в сантиметрах")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес в киллограммах")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight / 100, IMTPower)
	fmt.Printf("Индекс массы тела: %.0f", IMT)
}
