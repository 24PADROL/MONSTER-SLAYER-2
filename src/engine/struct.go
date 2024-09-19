
package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
)

type engine int

const (
	INGAME  engine = iota
	PAUSE    engine = iota
	INVENTORY engine = iota
	GAMEOVER engine = iota
	FIGHT 	engine = iota
	COFFRE engine = iota
	SHOP engine = iota
	ChaseDistance = 100

)

type Engine struct {
	Player   entity.Player
	Monsters []entity.Monster
	Shop  []entity.Shop
	Coffre	[]entity.Coffre
	Fight Fight
	
	Music       rl.Music
	MusicVolume float32

	Sprites map[string]rl.Texture2D

	Camera rl.Camera2D

	MapJSON MapJSON

	LoadingScreen rl.Texture2D
	LoadingScreenCombat rl.Texture2D
	LoadingScreenGameOver rl.Texture2D
	LoadingScreenPause rl.Texture2D

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
}

type Fight struct {
	CurrentMonster entity.Monster
}
