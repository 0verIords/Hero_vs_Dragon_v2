package main

import "fmt"

type Persona struct {
	HP     int
	damage int
}

var Hero = Persona{100, 10}
var Dragon = Persona{100, 25}
var i int

func GameStatus() {
	i++
	fmt.Println("#+++++++++++++++++++++++++++++++")
	fmt.Println("#Ход: ", i)
	fmt.Println("#Здоровье Героя: ", Hero.HP)
	fmt.Println("#Здоровье дракона: ", Dragon.HP)
	fmt.Println("#+++++++++++++++++++++++++++++++")
}

func main() {

}
