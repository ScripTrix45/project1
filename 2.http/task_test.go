package main

import (
	"testing"
)

func TestHttpWorkers(t *testing.T) {
	m := make(map[int]struct{})

	res := HttpWorkers()

	if len(res) != 500 {
		t.Errorf("Ожидалось 500 значений, получено: %v", len(res))
	}

	for i := 0; i < len(res); i++ {
		m[res[i].Id] = struct{}{}
	}

	if len(m) != 500 {
		t.Errorf("Ожидалось 500 уникальных значений ID, получено: %v", len(m))
	}

}
