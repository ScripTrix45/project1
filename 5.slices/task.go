package main

import "fmt"

func sliceTask(sl1, sl2 []int) []int {

	m := make(map[int]struct{})
	resSl := make([]int, 0)

	for _, val := range sl2 {
		m[val] = struct{}{}
	}

	fmt.Println(m)

	for _, val := range sl1 {
		if _, exists := m[val]; !exists {
			resSl = append(resSl, val)
		}
	}
	return resSl
}

func main() {
	sl1 := []int{2, 4, 10, 11, 12, 15}
	sl2 := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	resSl := sliceTask(sl1, sl2)
	fmt.Println(resSl)

}
