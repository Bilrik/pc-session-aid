package class

type Class struct {
	Name   string
	hitDie uint
	level  uint
}

var (
	Barbarian = Class{
		Name:   "Barbarian",
		hitDie: 12,
	}

	Bard = Class{
		Name:   "Bard",
		hitDie: 6,
	}

	Cleric = Class{
		Name:   "Cleric",
		hitDie: 8,
	}

	Druid = Class{
		Name:   "Druid",
		hitDie: 8,
	}

	Fighter = Class{
		Name:   "Fighter",
		hitDie: 10,
	}

	Monk = Class{
		Name:   "Monk",
		hitDie: 8,
	}

	Paladin = Class{
		Name:   "Paladin",
		hitDie: 10,
	}

	Ranger = Class{
		Name:   "Ranger",
		hitDie: 8,
	}

	Rogue = Class{
		Name:   "Rogue",
		hitDie: 6,
	}

	Sorcerer = Class{
		Name:   "Sorcerer",
		hitDie: 4,
	}

	Wizard = Class{
		Name:   "Wizard",
		hitDie: 4,
	}
)
