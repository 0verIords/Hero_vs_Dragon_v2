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
	var out = []int{100, 75, 100, 100, 55, 100}
	for i := 0; i < len(out); i++ {
		heroHill(&in[i])
		if in[i] != out[i] {
			t.Errorf("Неправильный выход! (%v) != (%v)", in[i], out[i])
		}

	}
}

// func TestGameStatus(t *testing.T) {
// 	game := Game(15, 55, 44)
// 	s := strings.Repeat("#", 15)
// 	var in = gameStatus(game)
// 	var out = fmt.Sprintf("%s \n#Ход: %v \n#Здоровье Героя: %v \n#Здоровье дракона: %v \n%s", s, 15, 55, 44, s)
// 	if in != out {
// 		t.Errorf("Неправильная проверка строк")
// 	}
// }
