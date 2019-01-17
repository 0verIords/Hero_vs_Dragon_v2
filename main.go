package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Persona struct {
	HP     int
	damage int
}

var Hero = Persona{100, 10}
var Dragon = Persona{100, 25}

func GameStatus(i int) int {
	fmt.Println("#+++++++++++++++++++++++++++++++")
	fmt.Println("#Ход: ", i)
	fmt.Println("#Здоровье Героя: ", Hero.HP)
	fmt.Println("#Здоровье дракона: ", Dragon.HP)
	fmt.Println("#+++++++++++++++++++++++++++++++")
	i++
	return i
}

func heroTurn(damage, answer int) int {
	if damage != 0 {
		Hero.HP -= damage
	}
	if answer == 1 {

		miss := rand.Intn(100)
		if miss >= 50 {
			return Hero.damage
		} else {
			return 0
		}
	} else {
		if Hero.HP >= 100 {
			Hero.HP = 100
		} else {
			Hero.HP += 20
		}
		return 0
	}
}

func dragonTurn(damage int) int {
	if damage != 0 {
		Dragon.HP -= damage
	}

	miss := rand.Intn(100)
	if miss >= 60 {
		return Dragon.damage
	} else {
		return 0
	}
}

func GameRun() {
	var i int
	var answer int
	var damage int
	for {
		i = GameStatus(i)
		fmt.Println("#1) Ударить")
		fmt.Println("#2) Лечиться")
		fmt.Print("#Выбор: ")
		fmt.Fscan(os.Stdin, &answer)
		if answer >= 3 {
			i--
		} else {
			damage = heroTurn(damage, answer)
			damage = dragonTurn(damage)
		}
	}
}

func main() {
	GameRun()
}
