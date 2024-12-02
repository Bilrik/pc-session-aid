package race

import (
	class "github.com/CyTechNomad/pc-session-aid/pkg/Class"
	stats "github.com/CyTechNomad/pc-session-aid/pkg/Stats"
)

var Gnome = gnome{
	race{
		Name:         "Gnome",
		Size:         stats.Small,
		Speed:        20,
		favoredClass: class.Bard,
	},
}

type gnome struct {
	race
}

func (g gnome) GetName() string {
	return g.Name
}

func (g gnome) GetSize() stats.Size {
	return g.Size
}

func (g gnome) GetSpeed() uint {
	return g.Speed
}

func (g gnome) GetFavoredClass() class.Class {
	return g.favoredClass
}
