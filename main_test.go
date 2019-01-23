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
		answer := selectCheck(in[i])
		if answer != "1" && answer != "2" {
			if answer != correct_answer {
				t.Errorf("Не правильная работа выбора")
			}
		}
	}
}

func TestCheckWin(t *testing.T) {
	var in = []int{100, 0, -20, 10, 30}
	var out = []bool{true, false, false, true, true}
	for i := 0; i < len(in); i++ {
		game := Game{i, in[i], 20}
		answer := game.gameWin()
		if answer != out[i] {
			t.Error("Не правильно работает функция gameWin")
		}
	}
}

func TestHeroTurn(t *testing.T) {
	var in = []string{"1", "0", "2", "1", "2", "4", "5", "6"}
	for i := 0; i < len(in); i++ {
		hero := Hero{100, 10}
		answer := hero.heroTurn(in[i])
		if answer != 0 && answer != hero.Damage {
			t.Error("Не правильная работа heroTurn")
		}
	}
}

func TestDragonTurn(t *testing.T) {
	var in = []int{10, 40, 60, 100, 90, 6}
	for i := 0; i < len(in); i++ {
		answer := dragonTurn(in[i])
		if answer != 0 && answer != in[i] {
			t.Error("Не правильная работа dragonTurn")
		}
	}
}

func TestGameInput(t *testing.T) {
	var in = []string{"1", "0", "2", "1", "2", "4", "5", "6"}
	for i := 0; i < len(in); i++ {
		answer := gameInput(in[i])
		if answer != in[i] {
			t.Error("Не правильная работа gameInput")
		}
	}
}
