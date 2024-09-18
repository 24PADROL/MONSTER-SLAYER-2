// package fight

// import (
// 	"main/src/engine"
// 	"main/src/entity"

// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

// type fight int

// const (
// 	PLAYER_TURN  fight = iota
// 	MONSTER_TURN fight = iota
// )

// func Fight(e* player entity.Player, monster entity.Monster) {
	
// 	if e.Player.Health <= 0 {
// 		e.Player.IsAlive = false
// 		e.Player.Money /= 2
// 		e.StateEngine = GAMEOVER
// 	} else if e.Fight.CurrentMonster.Health <= 0 {
// 		e.Monsters = append(e.Monsters[:e.Fight.CurrentMonsterIndex], e.Monsters[e.Fight.CurrentMonsterIndex+1:]...)
// 		e.StateMenu = PLAY
// 		e.StateEngine = INGAME
// 		e.Player.Inventory = append(e.Player.Inventory, e.Fight.CurrentMonster.Loot...)
// 		e.Player.Money += e.Fight.CurrentMonster.Worth
// 		// fmt.Println("----------------DEAD-------------------")
// 	} else {
// 		// fmt.Println("----------------COMBAT-------------------")
// 		if rl.IsKeyPressed(rl.KeyE) {
// 			e.Player.Attack(&e.Fight.CurrentMonster)
// 			e.Fight.CurrentMonster.Attack(&e.Player)
// 		}

		

// 		// e.Player.ToString()
// 		// e.Fight.CurrentMonster.ToString()
		
// 	}
// }