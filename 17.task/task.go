package main

import (
	"fmt"
	"time"
)

func binaryFound(sl, findValues []int) []int {
	var checkedNum int

	resSl := make([]int, 0)

	for _, val := range findValues {
		leftBorder := 0
		rightBorder := len(sl) - 1

		for {

			if leftBorder > rightBorder {
				fmt.Println("Значение", val, "не найдено")
				break
			}

			checkedNum = (rightBorder + leftBorder) / 2
			fmt.Println("Левая граница:", leftBorder, "Правая граница:", rightBorder)
			time.Sleep(1 * time.Second)
			fmt.Println("Будем проверять индекс:", checkedNum)
			if sl[checkedNum] == val {
				fmt.Println("Найдено совпадение:", val, sl[checkedNum])
				resSl = append(resSl, val)
				break
			} else if sl[checkedNum] > val {
				time.Sleep(1 * time.Second)
				rightBorder = checkedNum - 1
				fmt.Println("Взято число больше необходимого", sl[checkedNum], "новая правая граница:", rightBorder)
			} else if sl[checkedNum] < val {
				time.Sleep(1 * time.Second)
				leftBorder = checkedNum + 1
				fmt.Println("Взято число меньше необходимого", sl[checkedNum], "новая левая граница:", leftBorder)
			}
		}
	}
	return resSl
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	findValues := []int{1, 4, 9, 10, 11}
	resSl := binaryFound(sl, findValues)
	fmt.Println(resSl)
}
