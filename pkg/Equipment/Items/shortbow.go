package items

import equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"

var Shortbow = shortbow{
	Item: equipment.Item{
		Name:        "Shortbow",
		Description: "A shortbow",
		Weight:      2.0,
		Cost:        25.0,
	},
	requiredAmmoType: "Arrow",
}

type shortbow struct {
	equipment.Item
	requiredAmmoType string
}

func (s shortbow) GetDamage() uint {
	return 1
}

func (s shortbow) RequiredAmmo() (string, bool) {
	return s.requiredAmmoType, true
}
