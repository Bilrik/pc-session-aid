package items

import equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"

var Luckstone = luckstone{
	Item: equipment.Item{
		Name:        "Luckstone",
		Description: "A luckstone",
		Weight:      1.0,
		Cost:        100.0,
	},
}

type luckstone struct {
	equipment.Item
}

func (l luckstone) GetDamage() uint {
	return 1
}

func (l luckstone) GetACBonus() int {
	return 1
}
