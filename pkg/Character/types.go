package character

import (
	class "github.com/CyTechNomad/pc-session-aid/pkg/Class"
	equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"
	race "github.com/CyTechNomad/pc-session-aid/pkg/Race"
	stats "github.com/CyTechNomad/pc-session-aid/pkg/Stats"
)

type Character struct {
	name         string
	age          uint
	height       string
	weight       uint
	hp           stats.HP
	strength     stats.Ability
	dexterity    stats.Ability
	constitution stats.Ability
	intelligence stats.Ability
	wisdom       stats.Ability
	charisma     stats.Ability
	race         race.Race
	class        class.Class
	inventory    map[equipment.Equipment]int
	primaryHand  equipment.Equipment
}
