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
	for {
		usrHeight, usrKg := getUserInput()
		IMT := calculateIMT(usrKg, usrHeight)
		outputResult(IMT)
		isRepeateCalc := checkRepeatCalculation()
		if isRepeateCalc {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Индекс массы тела: %.0f", imt)
	fmt.Print(result)
	switch {
	case imt < 16:
		fmt.Println("\nУ вас сильный дефицит массы тела!")
	case imt < 18.5:
		fmt.Println("\nУ вас дефицит массы тела!")
	case imt < 25:
		fmt.Println("\nУ вас нормальный вес!")
	case imt < 30:
		fmt.Println("\nУ вас избыточный вес!")
	default:
		fmt.Println("\nТы жирдяй!")
	}
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

func checkRepeatCalculation() bool {
	var answer string
	fmt.Println("\nХотите сделать расчет еще раз?(да/нет)\n*****************************")
	fmt.Scan(&answer)
	return answer == "нет"
}
