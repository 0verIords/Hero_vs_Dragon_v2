package main

import (
	"testing"
)

func TestDamage(t *testing.T) {
	var in = []int{50, 40, 100, 90}
	var out = []int{50, 60, 0, 10}
	for i := 0; i < len(out); i++ {
		hitPoint := 100
		hitPoint -= hit(in[i])
		if hitPoint != out[i] {
			t.Errorf("Неправильный выход! (%v) (%v)", hitPoint, out[i])
		}

	}
}

func TestHill(t *testing.T) {
	var in = []int{100, 55, 81, 90, 35, 96}
	var out = []int{0, 20, 19, 10, 20, 4}
	for i := 0; i < len(out); i++ {
		hill := heroHill(in[i])
		if hill != out[i] {
			t.Errorf("Неправильный выход! (%v) (%v)", hill, out[i])
		}

	}
}
