package main

import (
	"fmt"
	"sort"
)

func deleteDoubles(s []int) []int {

	m := make(map[int]struct{})
	resSl := make([]int, 0)

	for _, val := range s {
		m[val] = struct{}{}
	}
	for key, _ := range m {
		resSl = append(resSl, key)
	}

	sort.Slice(resSl, func(i, j int) bool {
		return resSl[i] < resSl[j]
	})

	return resSl
}
func main() {
	sl := []int{0, 0, 1, 1, 1, 3, 3, 4, 5, 6, 7, 10}
	editedSl := deleteDoubles(sl)
	fmt.Println(editedSl)
}
