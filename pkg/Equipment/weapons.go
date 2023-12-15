package equipment

const (
	melee uint = 5
	reach uint = melee * 2
)

var (
	ShortSword = Weapon{
		Equipment: Equipment{
			Name:        "Short Sword",
			Description: "A short sword",
			Weight:      2.0,
			Cost:        10.0,
		},
		DamageDice:         6,
		Critical:           19,
		CriticalMultiplier: 2,
		Type:               "Piercing",
		Range:              melee,
	}
	ShortBow = Weapon{
		Equipment: Equipment{
			Name:        "Short Bow",
			Description: "A short bow",
			Weight:      2.0,
			Cost:        30.0,
		},
		DamageDice:         6,
		Critical:           20,
		CriticalMultiplier: 3,
		Type:               "Piercing",
		Range:              60,
	}
)
