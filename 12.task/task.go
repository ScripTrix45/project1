package main

import "fmt"

type Item struct {
	Category string
	Value    int
}

func intoMap(items []Item) map[string][]int {
	resultMap := make(map[string][]int)

	for _, item := range items {
		if values, exists := resultMap[item.Category]; !exists { // если ключ новый, создаем его в мапе и добавляем первое значение
			sl := []int{item.Value}
			resultMap[item.Category] = sl
		} else { // если ключ уже есть, то аппендим в слайс и записываем в мапу
			values = append(values, item.Value)
			resultMap[item.Category] = values
		}

	}
	return resultMap

}

func intoMap2(items []Item) map[string][]int {
	resultMap := make(map[string][]int)

	for _, item := range items {
		resultMap[item.Category] = append(resultMap[item.Category], item.Value)
	}

	return resultMap
}

func main() {
	sl := []Item{
		{Category: "A", Value: 10},
		{Category: "B", Value: 20},
		{Category: "A", Value: 30},
		{Category: "C", Value: 40},
		{Category: "B", Value: 50},
	}
	intoMap(sl)
	res := intoMap(sl)
	res1 := intoMap2(sl)
	fmt.Println(res, res1)
}
