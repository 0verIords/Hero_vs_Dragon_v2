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

func selectCheck(answer string) string {
	if answer != "1" && answer != "2" {
		return "#Неправильный ввод!\n#Нужно вводить только 1 или 2"
	} else {
		return answer
	}
}

func (g Game) gameWin() bool {
	if g.HeroHitPoint <= 0 {
		fmt.Printf("\n#Дракон выйграл за %v ходов\n", g.Counter)
		return false
	}
	if g.DragonHitPoint <= 0 {
		fmt.Printf("\n#Герой выйграл за %v ходов\n", g.Counter)
		return false
	}
	return true
}

func gameInput(answer string) string {

	fmt.Println("#1) Ударить")
	fmt.Println("#2) Лечиться")
	fmt.Print("#Выбор: ")
	fmt.Fscan(os.Stdin, &answer)
	return answer
}

func (hero *Hero) heroTurn(answer string) int {
	if answer == "1" {
		damage := miss(hit(hero.Damage))
		fmt.Printf("#Герой нанес %v урона\n", damage)
		return damage
	} else if answer == "2" {
		hero.heroHill()
		fmt.Println("#Герой вылечился")
		return 0
	} else {
		return 0
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

func dragonTurn(damage int) int {
	damage = miss(hit(damage))
	fmt.Printf("#Дракон нанес %v урона\n", damage)
	return damage
}

func gameRun(gameover bool) bool {
	var answer string
	var turn int

	hero := Hero{100, 10}
	dragon := Dragon{100, 40}
	for gameover {
		game := Game{turn, hero.HitPoint, dragon.HitPoint}
		gameover = game.gameWin()
		if gameover {
			fmt.Println(game.gameStatus())
			answer = gameInput(answer)
			answer = selectCheck(answer)
			if answer != "1" && answer != "2" {
				fmt.Println(answer)
			} else {
				dragon.HitPoint -= hero.heroTurn(answer)
				hero.HitPoint -= dragonTurn(dragon.Damage)
				turn++
			}
		}
	}
	return true

}

func main() {
	gameRun(true)
}
