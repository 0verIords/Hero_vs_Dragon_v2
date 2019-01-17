package main

import (
	"fmt"
	"math/rand"
)

type Persona struct {
	HP     int
	damage int
}

var Hero = Persona{100, 10}
var Dragon = Persona{100, 25}

func GameStatus() {
	fmt.Println("#+++++++++++++++++++++++++++++++")
	fmt.Println("#Ход: ", i)
	fmt.Println("#Здоровье Героя: ", Hero.HP)
	fmt.Println("#Здоровье дракона: ", Dragon.HP)
	fmt.Println("#+++++++++++++++++++++++++++++++")
}

func heroTurn(damage, answer) int {
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
	}
}

func GameRun() {

	GameStatus()
}

func main() {
	GameRun()
}
