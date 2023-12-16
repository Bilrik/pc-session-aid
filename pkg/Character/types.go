package character

import (
	"errors"

	class "github.com/Bilrik/pc-session-aid/pkg/Class"
	race "github.com/Bilrik/pc-session-aid/pkg/Race"
	stats "github.com/Bilrik/pc-session-aid/pkg/Stats"
)

type Character struct {
	Name         string
	Age          uint
	Height       string
	Weight       uint
	HP           stats.HP
	Strength     stats.Ability
	Dexterity    stats.Ability
	Constitution stats.Ability
	Intelligence stats.Ability
	Wisdom       stats.Ability
	Charisma     stats.Ability
	Race         race.Race
	Class        class.Class
	inventory    []interface{}
}

var (
	ErrInvalidHP = errors.New("HP invalid")
)
