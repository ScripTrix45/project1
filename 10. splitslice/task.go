package main

import "fmt"

func splitSlice(s []int, n int) [][]int {
	var res [][]int
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		res = append(res, s[i:end])
	}
	return res
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	n := 11
	resSl := splitSlice(sl, n)
	fmt.Println(resSl)

}
