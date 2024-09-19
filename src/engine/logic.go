package engine

import (
	"main/src/entity"
	"main/src/item"

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


func (e *Engine) Consumable(Item item.Item) {
	if Item.IsConsumable {
		e.Player.Health += 5
		if e.Player.Health > e.Player.MaxHealth {
			e.Player.Health = e.Player.MaxHealth
		}
	}
}
func (e *Engine) OverLogic() {
	if e.Player.Health <= 0 {
		e.StateEngine = GAMEOVER
	}
}

func (e *Engine) FightLogic() { //[ 0, 5, 8, 6] [ 0, 8, 6]

	e.Battle()
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

func (e *Engine) CoffreCollisions() {
	for _, Coffre := range e.Coffre {
		if Coffre.Position.X > e.Player.Position.X-30 &&
			Coffre.Position.X < e.Player.Position.X+10 &&
			Coffre.Position.Y > e.Player.Position.Y-30 &&
			Coffre.Position.Y < e.Player.Position.Y+10 {
			if Coffre.Name == "Potion" {
				if rl.IsKeyPressed(rl.KeyE) {
					e.StateEngine = COFFRE
				}
			}
		}
	}
}

// func (e *Engine) ShopCollisions() {
// 	for _, Shop := range e.Shop {
// 		if Shop.Position.X > e.Player.Position.X-30 &&
// 			Shop.Position.X < e.Player.Position.X+10 &&
// 			Shop.Position.Y > e.Player.Position.Y-30 &&
// 			Shop.Position.Y < e.Player.Position.Y+10 {
// 			if Shop.Name == "Potion" {
// 				if rl.IsKeyPressed(rl.KeyE) {
// 					e.StateEngine = COFFRE
// 				}
// 			}
// 		}
// 	}
// }

func (e *Engine) CoffreLogic() {
	if rl.IsKeyPressed(rl.KeyE) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
	}
}

func (e *Engine) ShopLogic() {
	if rl.IsKeyPressed(rl.KeyE) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
	}
}

func (e *Engine) InGameLogic() {
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
		// e.Player.Animation = entity.WALK
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
		if rl.IsKeyPressed(rl.KeyM) {
			e.Player.Attack(&e.Fight.CurrentMonster)
			e.Fight.CurrentMonster.Attack(&e.Player)
		}
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
// func (e *Engine) CheckCollisionsWithObjects() bool {
// 	playerRectangle := rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 40, 40)
// 	// * 2 - 16
// 	for _, obj := range e.Objects {
// 		objectsRectangle := rl.NewRectangle(obj.X, obj.Y, obj.Width, obj.Height)
// 		if rl.CheckCollision(playerRectangle, objectsRectangle) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (e *Engine) BlockCollisions(){
// 	if e.CheckCollisionsWithObjects(){
// 		if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
// 			e.Player.Position.X -= e.Player.Speed
// 		}
// 		if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
// 			e.Player.Position.Y += e.Player.Speed
// 		}
// 		if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
// 			e.Player.Position.Y -= e.Player.Speed
// 		}
// 	}
// }
func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()
	e.TrackMonsterLogic()
	e.CoffreCollisions()
}

// func (e *Engine) CheckCollisionsWithObjects() bool {
// 	playerRectangle := rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 40, 40)
// 	// * 2 - 16
// 	for _, obj := range e.Objects {
// 		objectsRectangle := rl.NewRectangle(obj.X, obj.Y, obj.Width, obj.Height)
// 		if rl.CheckCollision(playerRectangle, objectsRectangle) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (e *Engine) BlockCollisions() {
// 	if e.CheckCollisionsWithObjects() {
// 		if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
// 			e.Player.Position.X += e.Player.Speed
// 		}
// 		if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
// 			e.Player.Position.X -= e.Player.Speed
// 		}
// 		if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
// 			e.Player.Position.Y += e.Player.Speed
// 		}
// 		if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
// 			e.Player.Position.Y -= e.Player.Speed
// 		}
// 	}
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
