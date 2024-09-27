package engine

import (
	// "main/src/engine"

	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type fight int

const (
	PLAYER_TURN  fight = iota
	MONSTER_TURN fight = iota
)

func (e *Engine) Battle() {

	if e.Player.Health <= 0 {
		e.Player.IsAlive = false
		e.StateEngine = GAMEOVER
	} else if e.Fight.CurrentMonster.Health <= 0 {
		e.Monsters = append(e.Monsters[:e.Fight.CurrentMonster.Index], e.Monsters[e.Fight.CurrentMonster.Index+1:]...)
		e.StateMenu = PLAY
		e.StateEngine = INGAME
	
		fmt.Println(e.Player.Inventory)

		e.Player.Money += e.Fight.CurrentMonster.Worth
	} else {
		if rl.IsKeyPressed(rl.KeyE) {
			e.Player.Attack(&e.Fight.CurrentMonster)
			e.Fight.CurrentMonster.Attack(&e.Player)
		}

	}
}
