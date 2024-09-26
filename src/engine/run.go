package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Run() {
	rl.SetTargetFPS(120)

	for engine.IsRunning {

		rl.BeginDrawing()

		switch engine.StateMenu {
		case HOME:
			engine.HomeRendering()
			engine.HomeLogic()

		case SETTINGS:
			engine.SettingsLogic()
			engine.SettingsRendering()

		case PLAY:
			switch engine.StateEngine {
			case INGAME:
				engine.InGameRendering()
				engine.InGameLogic()

			case PAUSE:
				engine.PauseRendering()
				engine.PauseLogic()

			case FIGHT:
				engine.FightRendering()
				engine.FightLogic()

			case GAMEOVER:
				engine.OverRendering()
				engine.OverLogic()

			case COFFRE:
				engine.CoffreRendering()
				engine.CoffreLogic()

			case SHOP:
				engine.ShopRendering()
				engine.ShopLogic()
			}
		}

		rl.EndDrawing()
	}
}
