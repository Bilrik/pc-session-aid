package items

import equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"

var Arrow = arrow{
	Item: equipment.Item{
		Name:        "Arrow",
		Description: "An arrow",
		Weight:      0.1,
		Cost:        0.05,
	},
}

type arrow struct {
	equipment.Item
}

func (a arrow) GetAmmoType() string {
	return "Arrow"
}
