package engine

import (
	// "main/src/engine"

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
	e.Player.Money /= 2
	e.StateEngine = GAMEOVER
} else if e.Fight.CurrentMonster.Health <= 0 {
	e.Monsters = append(e.Monsters[:e.Fight.CurrentMonster.Index], e.Monsters[e.Fight.CurrentMonster.Index+1:]...)
	e.StateMenu = PLAY
	e.StateEngine = INGAME
	e.Player.Inventory = append(e.Player.Inventory, e.Fight.CurrentMonster.Loot...)
	e.Player.Money += e.Fight.CurrentMonster.Worth
	// fmt.Println("----------------DEAD-------------------")
} else {
	// fmt.Println("----------------COMBAT-------------------")
	if rl.IsKeyPressed(rl.KeyE) {
		e.Player.Attack(&e.Fight.CurrentMonster)
		e.Fight.CurrentMonster.Attack(&e.Player)
	}

	// e.e.Player.ToString()
	// e.Fight.CurrentMonster.ToString()

}
}
	
// 	if e.e.Player.Health <= 0 {
// 		e.e.Player.IsAlive = false
// 		e.e.Player.Money /= 2
// 		e.StateEngine = GAMEOVER
// 	} else if e.Fight.CurrentMonster.Health <= 0 {
// 		e.Monsters = append(e.Monsters[:e.Fight.CurrentMonsterIndex], e.Monsters[e.Fight.CurrentMonsterIndex+1:]...)
// 		e.StateMenu = PLAY
// 		e.StateEngine = INGAME
// 		e.e.Player.Inventory = append(e.e.Player.Inventory, e.Fight.CurrentMonster.Loot...)
// 		e.e.Player.Money += e.Fight.CurrentMonster.Worth
// 		// fmt.Println("----------------DEAD-------------------")
// 	} else {
// 		// fmt.Println("----------------COMBAT-------------------")
// 		if rl.IsKeyPressed(rl.KeyE) {
// 			e.e.Player.Attack(&e.Fight.CurrentMonster)
// 			e.Fight.CurrentMonster.Attack(&e.e.Player)
// 		}

		

// 		// e.e.Player.ToString()
// 		// e.Fight.CurrentMonster.ToString()
		
// 	}
// }

