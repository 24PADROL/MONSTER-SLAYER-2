package entity

import (
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shop struct {
	Name     string
	Position rl.Vector2
	Loot     []item.Item

	Sprite rl.Texture2D
}
