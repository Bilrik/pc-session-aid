package items

import equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"

var Torch = torch{
	equipment.Item{
		Name:        "Torch",
		Description: "A torch",
		Weight:      1.0,
		Cost:        1.0,
	},
}

type torch struct {
	equipment.Item
}
