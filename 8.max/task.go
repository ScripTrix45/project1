package main

import (
	"fmt"
	"slices"
)

func minMax(s []float64) (float64, float64) {
	var min float64
	var max float64

	for _, val := range s {
		if max < val {
			max = val
			fmt.Println("Новый максимум:", max)
		}
		if min > val {
			min = val
			fmt.Println("Новый минимум:", min)
		}
	}
	return min, max
}

func minMax1(s []float64) (float64, float64) {
	min := slices.Min(s)
	max := slices.Max(s)
	return min, max
}

func main() {
	sl := []float64{4.45, 1.96, 5.3, -145.34, 185.12}
	fmt.Println(sl)
	min, max := minMax(sl)
	fmt.Println(min, max)

	min1, max1 := minMax1(sl)
	fmt.Println(min1, max1)
}
