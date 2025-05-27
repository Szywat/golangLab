package main

import (
	"math"
	"stockmarket/lab06/structures"
)

func EMA(data []structures.Data) float64 {
	n := len(data)
	alfa := 2 / (n + 1)

	var ema float64
	var numerator float64
	var denominator float64

	for i := range n {
		numerator = numerator + data[n-i-1].Last*math.Pow(float64(1-alfa), float64(i))
		denominator = denominator + math.Pow(float64(1-alfa), float64(i))
	}

	ema = numerator / denominator

	return ema

}
