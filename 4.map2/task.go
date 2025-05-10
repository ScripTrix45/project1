package main

import (
	"fmt"
)

func mapTask(m map[int]struct{}) map[int]struct{} {

	firstNum := 0
	var counter int

	for val := range m {
		if counter == 0 {
			firstNum = val
			fmt.Println("Первое значение:", firstNum)
			counter++
		} else if counter == 1 {
			fmt.Println("Второе значение:", val)
			resNum := firstNum + val
			fmt.Println("результат:", resNum)
			m[resNum] = struct{}{}
			fmt.Println("Итоговая мапа:", m)
			counter++
		} else if counter == 2 {
			fmt.Println("Счетчик 2, прервано")
			break
		}
	}
	return m
}

func mapTask2(m map[int]struct{}) map[int]struct{} {

	keys := make([]int, 0, 2)
	for val := range m {
		keys = append(keys, val)
		if len(keys) == 2 {
			break
		}
	}
	resNum := keys[0] + keys[1]

	if _, exists := m[resNum]; !exists {
		m[resNum] = struct{}{}
	}
	return m
}

func main() {
	m := map[int]struct{}{
		1: struct{}{},
		2: struct{}{},
		3: struct{}{},
		8: struct{}{},
	}
	res := mapTask2(m)
	fmt.Println(res)
}
