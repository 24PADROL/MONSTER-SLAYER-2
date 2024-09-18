package item

import (
	"fmt"
)

type Potion struct {
	Name         string
	Price        int
	IsConsumable bool
	IsEquippable bool
}

func (i *Potion) ToString() {
	fmt.Printf("Je suis une potion qui vaut %d €\n", i.Price)
}
