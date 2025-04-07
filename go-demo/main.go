package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2
	var userHeight float64
	var userKg float64
	fmt.Println("Введите свой рост в метрах")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес в киллограммах")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print("Индекс массы тела: ")
	fmt.Print(IMT)
}
