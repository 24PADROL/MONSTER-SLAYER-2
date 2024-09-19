package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position  rl.Vector2
	Health    int 
	Money     int
	Speed     float32
	Inventory []item.Item
	Animation Animation
	MaxHealth int

	IsAlive bool

	Sprite rl.Texture2D
}

type Animation int

const (
	IDLE Animation = iota
	WALK Animation = iota
)

func (p *Player) Attack(m *Monster) {
	m.Health -= 3
}

func (p *Player) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		Inventaire: %+v
	
	\n`, p.Health, p.Money, p.Inventory)
}
