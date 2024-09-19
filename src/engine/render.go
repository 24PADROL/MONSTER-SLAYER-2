package engine

import (
	"main/src/entity"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue)
}

func (e *Engine) HomeRendering() {
	rl.ClearBackground(rl.Blue)
	rl.DrawTexture(e.LoadingScreen, 0, 0, rl.White)

	rl.DrawText("MONSTER SLAYER", int32(rl.GetScreenWidth())/2-rl.MeasureText("MONSTER SLAYER", 40)/2, int32(rl.GetScreenHeight())/2-200, 40, rl.RayWhite)
	rl.DrawText("2", int32(rl.GetScreenWidth())/2-rl.MeasureText("2", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[Enter] JOUER", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] JOUER", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Esc] QUITTER", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] QUITTER", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)
	rl.DrawText("[G] PARAMÈTRE", int32(rl.GetScreenWidth())/2-rl.MeasureText("[G] PARAMÈTRE", 20)/2, int32(rl.GetScreenHeight())/2+200, 20, rl.RayWhite)
	//Les settings marche pas alors qu'avant ça marchait
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()
	e.RenderCoffre()
	// e.RenderShop()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)

	rl.DrawText("Bienvenue !", int32(rl.GetScreenWidth())/2-rl.MeasureText("Bienvenue !", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] ou [Esc] pour mettre pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] ou [Esc] pour mettre en pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
	rl.DrawText("[Tab] pour ouvrir l'inventaire", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Tab] pour ouvrir l'inventaire", 20)/2, int32(rl.GetScreenHeight())/2-250, 20, rl.RayWhite)
	rl.DrawText("Pv :", int32(rl.GetScreenWidth())/23-rl.MeasureText("Pv :", 40)/2, int32(rl.GetScreenHeight())/2-375, 35, rl.Red)
	rl.DrawText("Argent :", int32(rl.GetScreenWidth())/19-rl.MeasureText("Argent :", 40)/3, int32(rl.GetScreenHeight())/2-325, 35, rl.Yellow)
	rl.DrawText(strconv.Itoa(e.Player.Health), int32(rl.GetScreenWidth())/7-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-375, 40, rl.Red)
	rl.DrawText(strconv.Itoa(e.Player.Money), int32(rl.GetScreenWidth())/6-rl.MeasureText("Home Menu", 40)/4, int32(rl.GetScreenHeight())/2-325, 40, rl.Yellow)
	rl.DrawText("FPS : " + strconv.Itoa(int(rl.GetFPS())), 20, 120, 40, rl.Green)
	rl.DrawRectangle(300, 650, 800, 100, rl.Purple)
	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(175, 500, 400, 400),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
		for _, monster := range e.Monsters {
			rl.DrawTexturePro(
				monster.Sprite,
				rl.NewRectangle(0, 0, 100, 100),
				rl.NewRectangle(350, 500, 400, 400),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White,
			)
		}
	
	
	//Affichage de la vie et de l'argent

	if rl.IsCursorOnScreen() {
		rl.HideCursor()
	}
}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(e.LoadingScreenPause, 0, 0, rl.White)
	rl.ClearBackground(rl.Red)

	rl.DrawText("Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[P] ou [Esc] pour continuer", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] ou [Esc] pour continuer", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour quitter", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

}

// JE TESTE L'OVERTURE DE L'INVENTAIRE
func (e *Engine) FightRendering() {
	rl.DrawTexture(e.LoadingScreenCombat, 0, 0, rl.White)
	rl.DrawText("Pv :", int32(rl.GetScreenWidth())/23-rl.MeasureText("Pv :", 40)/2, int32(rl.GetScreenHeight())/2-375, 35, rl.Red)
	rl.DrawText(strconv.Itoa(e.Player.Health), int32(rl.GetScreenWidth())/7-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-375, 40, rl.Red)
	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(100, 100, 600, 600),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

	for _, monster := range e.Monsters {
		rl.DrawTexturePro(
			monster.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(750, 100, 600, 600),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}

	rl.DrawText("COMBAT !!", int32(rl.GetScreenWidth())/2-rl.MeasureText("COMBAT !!", 40)/2, int32(rl.GetScreenHeight())/2-250, 40, rl.RayWhite)
}

func (e *Engine) OverRendering() {
	rl.DrawTexture(e.LoadingScreenGameOver, 0, 0, rl.White)
	rl.ClearBackground(rl.DarkGray)

	rl.DrawText("GAMEOVER HAHA ÇA MARCHE", int32(rl.GetScreenWidth())/2-rl.MeasureText("GAMEOVER HAHA ÇA MARCHE", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
}

func (e *Engine) CoffreRendering() {

	rl.ClearBackground(rl.Pink)
	rl.DrawText("Le coffre est vide pour l'instant tkt ça arrive", int32(rl.GetScreenWidth())/2-rl.MeasureText("Le coffre est vide pour l'instant tkt ça arrive", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
}

func (e *Engine) SettingsRendering() {
	rl.ClearBackground(rl.Yellow)

	rl.DrawText("[Z] pour aller en haut", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Z] pour aller vers le haut", 20)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("Paramètre", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paramètre", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	//JE TESTE L'OVERTURE DE L'INVENTAIRE
}

func (e *Engine) ShopRendering() {

	rl.ClearBackground(rl.Green)
	rl.DrawText("Le Shop est vide pour l'instant tkt ça arrive", int32(rl.GetScreenWidth())/2-rl.MeasureText("Le Shop est vide pour l'instant tkt ça arrive", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
}

func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}

// func (e *Engine) RenderShop() {

// 	rl.DrawTexturePro(
// 		Shop.Sprite,
// 		rl.NewRectangle(0, 0, 100, 100),
// 		rl.NewRectangle(Shop.Position.X, Shop.Position.Y, 200, 200),
// 		rl.Vector2{X: 0, Y: 0},
// 		0,
// 		rl.White,
// 	)
// }

func (e *Engine) RenderCoffre() {
	for _, coffre := range e.Coffre {
		rl.DrawTexturePro(
			coffre.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(coffre.Position.X, coffre.Position.Y, 200, 200),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		rl.DrawTexturePro(
			monster.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

func (e *Engine) RenderAnimationMonster() {
}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
