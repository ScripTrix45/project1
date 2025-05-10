package main

import (
	"testing"
)

func TestSorter(t *testing.T) {
	slicesCh := make(chan []int)
	exitCh := make(chan struct{})
	go func() {
		sl1 := []int{1, 3, 5, 7, 9}
		sl2 := []int{0, 2, 4, 6, 8, 10}
		slicesCh <- sl1
		slicesCh <- sl2
		exitCh <- struct{}{}
	}()

	result := sorter(slicesCh, exitCh)

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, but got %d", expected[i], i, v)
		} // БЛОЧНЫЕ ТЕСТЫ В GO(!)
	}
}
