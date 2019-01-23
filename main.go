package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Hero struct {
	HitPoint int
	Damage   int
}

type Dragon struct {
	HitPoint int
	Damage   int
}

type Game struct {
	Counter        int
	HeroHitPoint   int
	DragonHitPoint int
}

func (g Game) gameStatus() string {
	s := strings.Repeat("#", 15)
	return fmt.Sprintf("%s \n#Ход: %v \n#Здоровье Героя: %v \n#Здоровье дракона: %v \n%s", s, g.Counter, g.HeroHitPoint, g.DragonHitPoint, s)
}

func hit(hit int) int {
	return hit
}

func (hero *Hero) heroHill() {
	if hero.HitPoint <= 80 {
		hero.HitPoint += 20

	} else {
		hill := 100 - hero.HitPoint
		hero.HitPoint += hill

	}
}

func selectCheak(answer string) string {
	if answer != "1" && answer != "2" {
		return "#Неправильный ввод!\n#Нужно вводить только 1 или 2"
	} else {
		return answer
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
	var turn int
	var damage int

	hero := Hero{100, 10}
	dragon := Dragon{100, 40}
	for {
		game := Game{turn, hero.HitPoint, dragon.HitPoint}
		fmt.Println(game.gameStatus())
		fmt.Println("#1) Ударить")
		fmt.Println("#2) Лечиться")
		fmt.Print("#Выбор: ")
		fmt.Fscan(os.Stdin, &answer)
		answer = selectCheak(answer)
		if answer != "1" && answer != "2" {
			fmt.Println(answer)
		} else {
			if answer == "1" {
				damage = miss(hit(hero.Damage))
				fmt.Printf("#Герой нанес %v урона\n", damage)
				dragon.HitPoint -= damage
			} else {
				hero.heroHill()
				fmt.Println("#Герой вылечился")
			}
			damage = miss(hit(dragon.Damage))
			fmt.Printf("#Дракон нанес %v урона\n", damage)
			hero.HitPoint -= damage
			turn++
		}
		if hero.HitPoint <= 0 {
			fmt.Printf("#Дракон выйграл за %v ходов", turn)
			break
		}
		if dragon.HitPoint <= 0 {
			fmt.Printf("#Герой выйграл за %v ходов", turn)
			break
		}
	}

}

func main() {
	gameRun()
}
