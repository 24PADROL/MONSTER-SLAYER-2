package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Monster struct {
	Name     string
	Position rl.Vector2
	Health   int
	Damage   int
	Speed	float32
	Loot     []item.Item
	Worth    int //valeur en argent quand tué
	Index	int

	IsAlive bool

	Sprite rl.Texture2D
	Sprites []rl.Texture2D

	IsAnimated  bool
	FrameWidth  int
	FrameHeight int
	MaxFrames   int
}

var CurrentFrame int
var FrameCount int
var speed int= 5

func (m Monster) UpdateAnimation() {
	if FrameCount >= speed {
		CurrentFrame++
		FrameCount = 0
	} else {
		FrameCount++
	}
	if CurrentFrame >= m.MaxFrames {
		CurrentFrame = 0
	}
}

func (m *Monster) Draw() {
	FrameRec := rl.Rectangle{
		X: float32(m.FrameWidth + m.FrameWidth*CurrentFrame), Y: 0,
		Width: float32(m.FrameWidth), Height: float32(m.FrameHeight),
	}
	position := rl.Rectangle{
		X: m.Position.X + 20, Y: m.Position.Y + 10,
		Width: float32(m.FrameWidth), Height: float32(m.FrameHeight),
	}
	rl.DrawTexturePro(m.Sprite, FrameRec, position, rl.Vector2{}, 0, rl.White)
}

func (m *Monster) Attack(p *Player) {
	p.Health -= 10
}

func (m *Monster) ToString() {
	fmt.Printf("Je suis un monstre avec %d points de vie\n", m.Health)
}
