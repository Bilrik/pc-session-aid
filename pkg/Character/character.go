package character

import race "github.com/Bilrik/pc-session-aid/pkg/Race"

func WithCharacterName(name string) func(*Character) {
	return func(c *Character) {
		c.Name = name
	}
}

func WithAge(age uint) func(*Character) {
	return func(c *Character) {
		c.Age = age
	}
}

func WithHeight(height string) func(*Character) {
	return func(c *Character) {
		c.Height = height
	}
}

func WithAbilityScores(intelligence, charisma, strength, dexterity, constitution, wisdom uint) func(*Character) {
	return func(c *Character) {
		c.Intelligence.SetAbilityScore(intelligence)
		c.Charisma.SetAbilityScore(charisma)
		c.Strength.SetAbilityScore(strength)
		c.Dexterity.SetAbilityScore(dexterity)
		c.Constitution.SetAbilityScore(constitution)
		c.Wisdom.SetAbilityScore(wisdom)
	}
}

func WithRace(r race.Race) func(*Character) {
	return func(c *Character) {
		c.Race = r
	}
}

func WithMaxHP(hpMax uint) func(*Character) {
	return func(c *Character) {
		c.HP.Max = hpMax
		c.HP.Current = int(hpMax)
	}
}

func NewCharacter(opts ...func(*Character)) *Character {
	c := new(Character)

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Character) GetSpeed() uint {
	return c.Race.Speed
}

func (c *Character) GetAC() uint {
	ac := 10 + c.Dexterity.GetModifier() + c.Race.Size.GetACBonus()

	if ac < 0 {
		return 0
	}

	return uint(ac)
}
