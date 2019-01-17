package main

import (
	"testing"
)

func TestingGame(t *testing.T) {
	t.Run("Test damage", func(t *testing.T) {
		damage := heroTurn(20, 1)
		if damage != 10 && damage != 0 {
			t.Error("Incorect Hero damage")
		}
		damage = dragonTurn(20)
		if damage != 40 && damage != 0 {
			t.Error("Incorect Dragon damage")
		}
	})
}
