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
	rl.DrawText("Home Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[Enter] to Play", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] to Play", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Esc] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)
	rl.DrawText("[G] to WRAAAAAH Settings", int32(rl.GetScreenWidth())/2-rl.MeasureText("[G] to WRAAAAAH Settings", 20)/2, int32(rl.GetScreenHeight())/2+200, 20, rl.RayWhite)
	//Les settings marche pas alors qu'avant ça marchait
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("Bienvenue !", int32(rl.GetScreenWidth())/2-rl.MeasureText("Bienvenue !", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] ou [Esc] pour mettre pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] ou [Esc] pour mettre en pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
	rl.DrawText("[Tab] pour ouvrir l'inventaire", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Tab] pour ouvrir l'inventaire", 20)/2, int32(rl.GetScreenHeight())/2-250, 20, rl.RayWhite)
	rl.DrawText(strconv.Itoa(e.Player.Health), int32(rl.GetScreenWidth())/9-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-450, 40, rl.Red)
	rl.DrawText(strconv.Itoa(e.Player.Money), int32(rl.GetScreenWidth())/7-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-400, 40, rl.Yellow)
	rl.DrawText("Pv :", int32(rl.GetScreenWidth())/35-rl.MeasureText("Pv :", 40)/2, int32(rl.GetScreenHeight())/2-450, 35, rl.Red)
	rl.DrawText("Argent :", int32(rl.GetScreenWidth())/20-rl.MeasureText("Argent :", 40)/2, int32(rl.GetScreenHeight())/2-400, 35, rl.Yellow)
	//Affichage de la vie et de l'argent

	if rl.IsCursorOnScreen(){	
		rl.SetMouseCursor(3)
	}
}

func (e *Engine) PauseRendering() {
	rl.ClearBackground(rl.Red)

	rl.DrawText("Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[P] ou [Esc] pour continuer", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] ou [Esc] pour continuer", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q]/[A] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour quitter", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

}
func (e *Engine) InventoryRendering() {
	rl.ClearBackground(rl.Blue)

	rl.DrawText("C'est l'inventaire ça", int32(rl.GetScreenWidth())/2-rl.MeasureText("C'est l'inventaire ça", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	//JE TESTE L'OVERTURE DE L'INVENTAIRE
}
func (e *Engine) FightRendering() {
	
}

func (e *Engine) SettingsRendering() {
	rl.ClearBackground(rl.Yellow)

	rl.DrawText("Paramètre", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paramètre", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	//JE TESTE L'OVERTURE DE L'INVENTAIRE
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
func (e *Engine) RenderShoper() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Shoper.Position.X, e.Shoper.Position.Y, 250, 250),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
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
