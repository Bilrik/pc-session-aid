package character

import (
	class "github.com/CyTechNomad/pc-session-aid/pkg/Class"
	equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"
	race "github.com/CyTechNomad/pc-session-aid/pkg/Race"
	stats "github.com/CyTechNomad/pc-session-aid/pkg/Stats"
)

func WithCharacterName(name string) func(*Character) {
	return func(c *Character) {
		c.name = name
	}
}

func WithAge(age uint) func(*Character) {
	return func(c *Character) {
		c.age = age
	}
}

func WithHeight(height string) func(*Character) {
	return func(c *Character) {
		c.height = height
	}
}

func WithAbilityScores(intelligence, charisma, strength, dexterity, constitution, wisdom uint) func(*Character) {
	return func(c *Character) {
		c.intelligence.SetAbilityScore(intelligence)
		c.charisma.SetAbilityScore(charisma)
		c.strength.SetAbilityScore(strength)
		c.dexterity.SetAbilityScore(dexterity)
		c.constitution.SetAbilityScore(constitution)
		c.wisdom.SetAbilityScore(wisdom)
	}
}

func WithRace(r race.Race) func(*Character) {
	return func(c *Character) {
		c.race = r
	}
}

func WithClass(class class.Class) func(*Character) {
	return func(c *Character) {
		c.class = class
		c.class.LevelUp()
	}
}

func NewCharacter(opts ...func(*Character)) (*Character, error) {
	c := new(Character)

	for _, opt := range opts {
		opt(c)
	}

	if c.class != nil {
		if hp := int(c.class.GetHitDie()) + c.constitution.GetModifier(); hp < 1 {
			return nil, stats.ErrInvalidHP
		} else {
			c.hp.Max = uint(hp)
		}
		c.hp.Current = int(c.hp.Max)
	}

	c.inventory = make(map[equipment.Equipment]int)

	return c, nil
}

// Damage the character for hp. If hp is negative, return an error.
func (c *Character) Damage(hp uint) {
	c.hp.ModifyCurrent(-int(hp))
}

// Heal the character for hp. If hp is negative, return an error.
func (c *Character) Heal(hp uint) {
	c.hp.ModifyCurrent(int(hp))
}
