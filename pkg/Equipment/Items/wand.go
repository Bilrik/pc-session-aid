package items

import equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"

var Wand = wand{
	Item: equipment.Item{
		Name:        "Wand",
		Description: "A wand",
		Weight:      1.0,
		Cost:        10.0,
	},
	charges: 10,
}

type wand struct {
	equipment.Item
	charges uint
}

func (w wand) GetDamage() uint {
	return 1
}

func (w wand) GetCharges() uint {
	return w.charges
}

func (w *wand) UseCharge() {
	if w.charges > 0 {
		w.charges--
	}
}
