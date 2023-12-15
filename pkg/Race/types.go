package race

import (
	class "github.com/Bilrik/pc-session-aid/pkg/Class"
	stats "github.com/Bilrik/pc-session-aid/pkg/Stats"
)

type Race struct {
	Name         string
	Size         stats.Size
	Speed        uint
	favoredClass class.Class
}

var (
	Human = Race{
		Name:  "Human",
		Size:  stats.Medium,
		Speed: 30,
	}

	Dwarf = Race{
		Name:         "Dwarf",
		Size:         stats.Medium,
		Speed:        20,
		favoredClass: class.Fighter,
	}

	Elf = Race{
		Name:         "Elf",
		Size:         stats.Medium,
		Speed:        30,
		favoredClass: class.Wizard,
	}

	Gnome = Race{
		Name:         "Gnome",
		Size:         stats.Small,
		Speed:        20,
		favoredClass: class.Bard,
	}

	HalfElf = Race{
		Name:  "Half-Elf",
		Size:  stats.Medium,
		Speed: 30,
	}

	HalfOrc = Race{
		Name:         "Half-Orc",
		Size:         stats.Medium,
		Speed:        30,
		favoredClass: class.Barbarian,
	}

	Halfling = Race{
		Name:         "Halfling",
		Size:         stats.Small,
		Speed:        20,
		favoredClass: class.Rogue,
	}
)
