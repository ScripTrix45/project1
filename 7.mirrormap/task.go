package main

import "fmt"

func mapWorker(m map[string]int) map[int][]string {
	resMap := make(map[int][]string)

	for key, value := range m {

		if value2, exists := resMap[value]; !exists {
			sl := []string{key}
			resMap[value] = sl
		} else {
			value2 = append(value2, key)
			resMap[value] = value2
		}
	}
	return resMap
}

func main() {
	m := map[string]int{
		"cherry":    1,
		"apple":     1,
		"banana":    2,
		"kiwi":      3,
		"potato":    2,
		"pineapple": 3,
	}
	resmap := mapWorker(m)
	fmt.Println(resmap)
}
