package main

import "fmt"

func filterSlice(s []int, filter func(s int) bool) []int {
	resSl := make([]int, 0)

	for _, val := range s {
		if filter(val) {
			resSl = append(resSl, val)
		}
	}
	return resSl
}

func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := filterSlice(sl, func(s int) bool {
		return s%2 == 0
	})
	fmt.Println(res)
}
