package fight

import (
	"fmt"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type fight int

const (
	PLAYER_TURN  fight = iota
	MONSTER_TURN fight = iota
)

func Fight(player entity.Player, monster entity.Monster) {

	for { // Boucle infinie
		// Check si le joueur ou le monstre est vaincu. Si c'est le cas, on sort de la boucle
		if player.Health <= 0 {
			player.IsAlive = false
			player.Money /= 2
			break
		} else if monster.Health <= 0 {
			player.Inventory = append(player.Inventory, monster.Loot...)
			player.Money += monster.Worth
			fmt.Println("----------------DEAD-------------------")
			break
		} else {
			fmt.Println("----------------COMBAT-------------------")

			if rl.IsKeyDown(rl.KeyE) {
				player.Attack(&monster)
			}

			monster.Attack(&player)

			player.ToString()
			monster.ToString()
			
		}
	}
}