package engine

import (
	"main/src/entity"
	// "main/src/fight"
	

	//"main/src/fight"

	rl "github.com/gen2brain/raylib-go/raylib"
)
func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/Hello.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)
	}
	if rl.IsKeyPressed(rl.KeyG) {
		e.StateMenu = SETTINGS
		//CA MARCHE PAAAS AVANT CA MARCHAIT
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyG) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}
func (e *Engine) InventoryLogic() {
	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
	}
}
func (e *Engine) OverLogic() {
	if e.Player.Health <= 0 {
		e.StateEngine = GAMEOVER
	}
}

func (e *Engine) FightLogic() { //[ 0, 5, 8, 6] [ 0, 8, 6]

	if e.Player.Health <= 0 {
		e.Player.IsAlive = false
		e.Player.Money /= 2
		e.StateEngine = GAMEOVER
	} else if e.Fight.CurrentMonster.Health <= 0 {
		e.Monsters = append(e.Monsters[:e.Fight.CurrentMonsterIndex], e.Monsters[e.Fight.CurrentMonsterIndex+1:]...)
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

		

		// e.Player.ToString()
		// e.Fight.CurrentMonster.ToString()
		
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

	e.Battle()
 }
}
func (e *Engine) TrackMonsterLogic() {
	for i := 0; i < len(e.Monsters); i++ {

		if e.Monsters[i].IsAlive {
			distance := rl.Vector2Distance(e.Player.Position, e.Monsters[i].Position)

			if distance <= ChaseDistance {
				direction := rl.Vector2Subtract(e.Player.Position, e.Monsters[i].Position)
				direction = rl.Vector2Normalize(direction)
				e.Monsters[i].Position = rl.Vector2Add(e.Monsters[i].Position, direction)
			}
		}
	}

	// fight.Fight()
}

func (e *Engine) TrackMonsterLogic(){
    for i := 0; i < len(e.Monsters); i++ { 

        if e.Monsters[i].IsAlive {
            distance := rl.Vector2Distance(e.Player.Position, e.Monsters[i].Position)

            if distance <= ChaseDistance {
                direction := rl.Vector2Subtract(e.Player.Position, e.Monsters[i].Position)
                direction = rl.Vector2Normalize(direction)
                e.Monsters[i].Position = rl.Vector2Add(e.Monsters[i].Position, direction)
            }
        }
    }
}
func (e *Engine) InGameLogic() {
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
	}
	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateEngine = INVENTORY
	}


	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/Ost2.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}
func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()
	e.TrackMonsterLogic()
}
// func (e *Engine) WorldCollisions() {
// type def struct PhysicObject {
//     unsigned int id;
//     Transform transform;
//     Collider collider;
//     Rigidbody rigidbody;
//     bool enabled;
// } PhysicObject;
// }

func (e *Engine) MonsterCollisions() {

	for i, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-20 &&
			monster.Position.X < e.Player.Position.X+20 &&
			monster.Position.Y > e.Player.Position.Y-20 &&
			monster.Position.Y < e.Player.Position.Y+20 {

			monster.Index = i

			e.Fight.CurrentMonster = e.Monsters[i]

			e.StateEngine = FIGHT

			if monster.Name == "Yann" {
				if rl.IsKeyDown(rl.KeyE) {
					e.NormalTalk(monster, "RWAAAAAAAAAAH")
					//lancer un combat ?
				}
			}
		} else {
			e.NormalTalk(monster, "Vas-y viens")
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyA) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}

//NOOOOOOOOOON
