package main

import (
	"fmt"
	"strings"
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
		testHero := Hero{in[i], 10}
		testHero.heroHill()
		if testHero.HitPoint != out[i] {
			t.Errorf("Неправильный выход! (%v) != (%v)", testHero.HitPoint, out[i])
		}

	}
}

func TestGameStatus(t *testing.T) {
	game := Game{15, 55, 44}
	s := strings.Repeat("#", 15)
	var in = game.gameStatus()
	var out = fmt.Sprintf("%s \n#Ход: %v \n#Здоровье Героя: %v \n#Здоровье дракона: %v \n%s", s, 15, 55, 44, s)
	if in != out {
		t.Errorf("Неправильная проверка строк")
	}
}

func TestMiss(t *testing.T) {
	var in = []int{10, 40, 60, 100, 90, 6}
	var out = []int{10, 40, 60, 100, 90, 6}
	for i := 0; i < len(out); i++ {
		ans := miss(in[i])
		if ans != out[i] && ans != 0 {
			t.Errorf("random does not work correct")
		}
	}
}

func TestChoice(t *testing.T) {
	var in = []string{"1", "2", "qweqwe", "sdhsbhsa", "12", "!$@", "sdsda2", "wq"}
	correct_answer := "#Неправильный ввод!\n#Нужно вводить только 1 или 2"
	for i := 0; i < len(in); i++ {
		answer := selectCheak(in[i])
		if answer != "1" && answer != "2" {
			if answer != correct_answer {
				t.Errorf("Не правильная работа выбора")
			}
		}
	}
}
