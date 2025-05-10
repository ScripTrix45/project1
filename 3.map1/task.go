package main

import (
	"fmt"
	"sort"
)

func mapTask(s ...[]int) []int {
	m := make(map[int]struct{})

	resSl := make([]int, 0)

	for sl, _ := range s {
		for _, val := range s[sl] { // проверка в мапе должна быть
			m[val] = struct{}{}
		}

	}
	for val := range m {
		resSl = append(resSl, val)
	}

	sort.Slice(resSl, func(i, j int) bool {
		return resSl[i] < resSl[j]
	})

	return resSl
}

func main() {
	sl := []int{1, 1, 1, 2, 2, 3, 5, 6, 7, 8, 9, 10}
	sl1 := []int{3, 4, 5}
	resSl := mapTask(sl, sl1)
	fmt.Println(resSl)
}
