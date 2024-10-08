package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.Player.Sprite = rl.LoadTexture("textures/entities/soldier/Soldier-Idle.png")
	e.LoadingScreen = rl.LoadTexture("textures/menu/Grotte.jpg")
	e.LoadingScreenCombat = rl.LoadTexture("textures/menu/FondMenuCombat2.jpg")
	e.LoadingScreenGameOver = rl.LoadTexture("textures/menu/MortScreen.jpg")
	e.LoadingScreenPause = rl.LoadTexture("textures/menu/PauseMenu.jpg")
}

func (e *Engine) LoadCharacter() {
	switch e.Player.Animation {
	case entity.IDLE:
		e.Player.Sprite = rl.LoadTexture("textures/entities/soldier/Soldier-Idle.png")
	case entity.WALK:
		e.Player.Sprite = rl.LoadTexture("textures/entities/soldier/Soldier-Walk.png")

	}

}

func (e *Engine) UnloadCharacter() {
	rl.UnloadTexture(e.Player.Sprite)
}

func (e *Engine) Unload() {
	// On libère les textures chargées, le joueur, la map, les monstres, etc...
	rl.UnloadTexture(e.Player.Sprite)

	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}
}
