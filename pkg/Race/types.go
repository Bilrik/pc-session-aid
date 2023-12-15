package race

import stats "github.com/Bilrik/pc-session-aid/pkg/Stats"

type Race struct {
	Name  string
	Size  stats.Size
	Speed uint
}

var (
	Elf = Race{
		Name:  "Elf",
		Size:  stats.Medium,
		Speed: 30,
	}

	Dwarf = Race{
		Name:  "Dwarf",
		Size:  stats.Medium,
		Speed: 20,
	}

	Gnome = Race{
		Name:  "Gnome",
		Size:  stats.Small,
		Speed: 20,
	}
)
