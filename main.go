package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type ParameterHero struct {
	HitPoint int
	Damage   int
}

type ParameterDragon struct {
	HitPoint int
	Damage   int
}

type Game struct {
	i              int
	HeroHitPoint   int
	DragonHitPoint int
}

func gameStatus(g Game) string {
	s := strings.Repeat("#", 15)
	return fmt.Sprintf("%s \n#Ход: %v \n#Здоровье Героя: %v \n#Здоровье дракона: %v \n%s", s, g.i, g.HeroHitPoint, g.DragonHitPoint, s)
}

func hit(hit int) int {
	return hit
}

func heroHill(heroHitPoint int) int {
	if heroHitPoint <= 80 {
		return 20
	} else {
		hill := 100 - heroHitPoint
		return hill
	}
}

func miss(damage int) int {
	miss := rand.Intn(100)
	if damage == 10 {
		if miss >= 50 {
			return damage
		} else {
			return 0
		}
	} else {
		if miss >= 60 {
			return damage
		} else {
			return 0
		}
	}
}

func gameRun() {
	var answer string
	var i int
	var damage int

	hero := ParameterHero{100, 10}
	dragon := ParameterDragon{100, 40}
	for {
		game := Game{i, hero.HitPoint, dragon.HitPoint}
		fmt.Println(gameStatus(game))
		fmt.Println("#1) Ударить")
		fmt.Println("#2) Лечиться")
		fmt.Print("#Выбор: ")
		fmt.Fscan(os.Stdin, &answer)
		if answer != "1" && answer != "2" {
			fmt.Println("#Неправильный ввод!")
			fmt.Println("#Нужно вводить только 1 или 2")
		} else {
			if answer == "1" {
				damage = miss(hit(hero.Damage))
				fmt.Printf("#Герой нанес %v урона\n", damage)
				dragon.HitPoint -= damage
			} else {
				hill := heroHill(hero.HitPoint)
				fmt.Printf("#Герой вылечился на %v хп\n", hill)
				hero.HitPoint += hill
			}
			damage = miss(hit(dragon.Damage))
			fmt.Printf("#Дракон нанес %v урона\n", damage)
			hero.HitPoint -= damage
			i++
		}
		if hero.HitPoint <= 0 {
			fmt.Printf("#Дракон выйграл за %v ходов", i)
			break
		}
		if dragon.HitPoint <= 0 {
			fmt.Printf("#Герой выйграл за %v ходов", i)
			break
		}
	}

}

func main() {
	gameRun()
}
