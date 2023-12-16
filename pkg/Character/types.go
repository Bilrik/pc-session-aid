package character

import (
	"errors"

	class "github.com/Bilrik/pc-session-aid/pkg/Class"
	race "github.com/Bilrik/pc-session-aid/pkg/Race"
	stats "github.com/Bilrik/pc-session-aid/pkg/Stats"
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
	inventory    []interface{}
}

var (
	ErrInvalidHP = errors.New("HP invalid")
)
