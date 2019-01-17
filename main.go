package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Parameter struct {
	HP     int
	damage int
}

var Hero = Parameter{100, 10}
var Dragon = Parameter{100, 40}

func GameStatus(i int) int {
	fmt.Println("#+++++++++++++++++++++++++++++++")
	fmt.Println("#Ход: ", i)
	fmt.Println("#Здоровье Героя: ", Hero.HP)
	fmt.Println("#Здоровье дракона: ", Dragon.HP)
	fmt.Println("#+++++++++++++++++++++++++++++++")
	i++
	return i
}

func HeroHit() {
	miss := rand.Intn(100)
	if miss >= 50 {
		Dragon.HP -= Hero.damage
	}
}

func HeroHealing() {
	Hero.HP += 20
	if Hero.HP >= 100 {
		Hero.HP = 100
	}
}

func HeroTurn(answer int) {
	if answer == 1 {
		HeroHit()
	} else {
		HeroHealing()
	}
}

func DragonTurn() {
	miss := rand.Intn(100)
	if miss >= 60 {
		Hero.HP -= Dragon.damage
	}
}

func GameRun() {
	var i int
	var answer int
	for {
		i = GameStatus(i)
		fmt.Println("#1) Ударить")
		fmt.Println("#2) Лечиться")
		fmt.Print("#Выбор: ")
		fmt.Fscan(os.Stdin, &answer)
		if answer != 1 && answer != 2 {
			i--
		} else {
			HeroTurn(answer)
			DragonTurn()
			if Hero.HP <= 0 {
				fmt.Println("Герой проиграл за ", i, " ходов")
				break
			}
			if Dragon.HP <= 0 {
				fmt.Println("Герой победил за ", i, " ходов")
				break
			}
		}
	}
}

func main() {
	GameRun()
}
