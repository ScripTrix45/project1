package main

import "fmt"

func mergeMap(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)
	for key, val := range m1 {
		result[key] = val
	}

	for key, val := range m2 {
		if val1, exists := result[key]; !exists || val > val1 { // если ключа из m2 нет в m1 или значение есть, но значение из m2 больше значения из m1
			result[key] = val // записать ключ = значение из m2
		}
	}
	return result
}

func main() {
	m1 := map[string]int{
		"Alex": 30,
		"Bob":  45,
		"John": 19,
	}
	m2 := map[string]int{
		"Daniel": 20,
		"Egor":   34,
		"Alex":   20,
		"Bob":    48,
	}
	res := mergeMap(m1, m2)
	fmt.Println(res)
}
