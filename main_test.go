package main

import (
	"testing"
)

func damageTest(t *testing.T) {
	var in = []int{50, 40, 100, 90}
	var out = []int{60, 60, 0, 10}
	for i := 0; i < len(out); i++ {
		hitPoint := 100
		hitPoint -= hit(in[i])
		if hitPoint != out[i] {
			t.Errorf("Неправильный выход! (%v) (%v)", hitPoint, out[i])
		}

	}
}

func hillTest(t *testing.T) {
	var in = []int{100, 55, 15, 90, 35}
	var out = []int{100, 75, 35, 100, 55}
	for i := 0; i < len(out); i++ {
		hill := heroHill(in[i])
		if hill != out[i] {
			t.Errorf("Неправильный выход! (%v) (%v)", hill, out[i])
		}

	}
}
