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

func HeroHit(damage int) int {
	miss := rand.Intn(100)
	if miss >= 50 {
		return damage
	} else {
		return 0
	}
}

func HeroHealing(HP int) int {
	HP += 20
	if HP >= 100 {
		return 100
	} else {
		return HP
	}
}

func HeroTurn(answer int) {
	if answer == 1 {
		Dragon.HP -= HeroHit(Hero.damage)
	} else {
		Hero.HP = HeroHealing(Hero.HP)
	}
}

func DragonTurn(damage int) int {
	miss := rand.Intn(100)
	if miss >= 60 {
		return damage
	} else {
		return 0
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
			Hero.HP -= DragonTurn(Dragon.damage)
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
