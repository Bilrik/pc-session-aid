package race

type Race struct {
	Name  string
	Size  string
	Speed uint
}

var (
	Elf = Race{
		Name:  "Elf",
		Size:  "Medium",
		Speed: 30,
	}

	Dwarf = Race{
		Name:  "Dwarf",
		Size:  "Medium",
		Speed: 20,
	}

	Gnome = Race{
		Name:  "Gnome",
		Size:  "Small",
		Speed: 20,
	}
)
