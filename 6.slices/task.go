package main

import "fmt"

func slicesTask(sl1, sl2 []string) []string {
	m := make(map[string]struct{})
	resSl := make([]string, 0)

	for _, val := range sl1 {
		m[val] = struct{}{}
	}

	for _, val := range sl2 {
		if _, exists := m[val]; exists {
			resSl = append(resSl, val)
			delete(m, val)
		}
	}
	return resSl
}

func main() {
	sl1 := []string{"first", "second", "third", "fourth", "fiveth", "sixth"}
	sl2 := []string{"first", "first", "first", "fiveth", "twelve"}
	resSl := slicesTask(sl1, sl2)
	fmt.Println(resSl)
}
