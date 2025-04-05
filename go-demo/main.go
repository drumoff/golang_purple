package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2
	userHeight := 1.8
	userKg := 100.0
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT) 
}