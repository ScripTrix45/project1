package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func sorter(c <-chan []int, e <-chan struct{}) []int {
	resSl := make([]int, 0, 1024)
	for {
		select {
		case val := <-c:
			{
				resSl = append(resSl, val...)
			}
		case <-e:
			sort.Slice(resSl, func(i, j int) bool {
				return resSl[i] < resSl[j]
			})
			return resSl
		}
	}
}

func worker(c chan<- []int) {
	for {
		capSl := rand.Intn(2) + 1
		sl := make([]int, 0, capSl)
		for len(sl) < cap(sl) {
			randomInt := rand.Intn(100) - 50
			sl = append(sl, randomInt)
		}
		fmt.Println("Входящий слайс:", sl)
		time.Sleep(1 * time.Second)
		c <- sl
	}
}

func exitFromSorter(e chan<- struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "exit" {
			e <- struct{}{}
		}

	}
}

func main() {
	slicesCh := make(chan []int)
	exitCh := make(chan struct{})
	go worker(slicesCh)
	go exitFromSorter(exitCh)
	result := sorter(slicesCh, exitCh)
	fmt.Println("Результат:", result)

}
