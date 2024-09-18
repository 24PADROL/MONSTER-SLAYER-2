package engine

import (
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 800
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Monster Slayer 2")

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")
}


func (e *Engine) InitEntities() {

	e.Coffre = append(e.Coffre, entity.Coffre{
		Name:     "Potion",
		Position: rl.Vector2{X: 600, Y: 320},
		Loot:     []item.Item{},

		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
		
	})
	

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 932, Y: 640},
		Health:    100,
		Money:     1000,
		Speed:     2,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Yann",
		Position: rl.Vector2{X: 932, Y: 320},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,
		Index:    len(e.Monsters),
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Player.Money = 120
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/Hello.mp3")

	rl.PlayMusicStream(e.Music)
}

//func (e *Engine) InitItem() {
	//Name:	"Seringue",
//}

//bonjourvvv


