package character

import (
	"fmt"
	"os"
	"text/tabwriter"

	class "github.com/Bilrik/pc-session-aid/pkg/Class"
	race "github.com/Bilrik/pc-session-aid/pkg/Race"
)

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

func WithClass(class class.Class) func(*Character) {
	return func(c *Character) {
		c.Class = class
		c.Class.LevelUp()
	}
}

func NewCharacter(opts ...func(*Character)) (*Character, error) {
	c := new(Character)

	for _, opt := range opts {
		opt(c)
	}

	if hp := int(c.Class.GetHitDie()) + c.Constitution.GetModifier(); hp < 1 {
		return nil, ErrInvalidHP
	} else {
		c.HP.Max = uint(hp)
	}
	c.HP.Current = int(c.HP.Max)

	return c, nil
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

func (c *Character) LevelUp(takeAverage bool) {
	c.Class.LevelUp()
	hpGain := c.Class.GetHPRoll()
	if takeAverage {
		hpGain = c.Class.GetHPAverage()
	}
	c.HP.ModifyMax(int(hpGain))
}

func (c *Character) GetFortitudeSave() int {
	return c.Constitution.GetModifier()
}

func (c *Character) GetReflexSave() int {
	return c.Dexterity.GetModifier()
}

func (c *Character) GetWillSave() int {
	return c.Wisdom.GetModifier()
}

func (c *Character) GetInventory() []interface{} {
	return c.equipment
}

func (c *Character) AddItem(i interface{}) {
	c.equipment = append(c.equipment, i)
}

func (c *Character) RemoveItem(i interface{}) {
	for index, item := range c.equipment {
		if item == i {
			c.equipment = append(c.equipment[:index], c.equipment[index+1:]...)
			return
		}
	}
}

func (c *Character) Print() {
	fmt.Printf("Name: %s \t Race: %s\n", c.Name, c.Race.Name)
	fmt.Printf("Age: %d \t Height: %s\t Weight: %d\n", c.Age, c.Height, c.Weight)
	fmt.Printf("Level: %d, Class: %s\n", c.Class.GetLevel(), c.Class.Name)

	fmt.Printf("HP: %d/%d\n", c.HP.Current, c.HP.Max)
	fmt.Printf("AC: %d\n", c.GetAC())
	fmt.Printf("Speed: %d\n", c.GetSpeed())

	fmt.Printf("Stats:\n")
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "\tStrength:\t", c.Strength.Score, "\t", c.Strength.GetModifier())
	fmt.Fprintln(writer, "\tDexterity:\t", c.Dexterity.Score, "\t", c.Dexterity.GetModifier())
	fmt.Fprintln(writer, "\tConstitution:\t", c.Constitution.Score, "\t", c.Constitution.GetModifier())
	fmt.Fprintln(writer, "\tWisdom:\t", c.Wisdom.Score, "\t", c.Wisdom.GetModifier())
	fmt.Fprintln(writer, "\tIntelligence:\t", c.Intelligence.Score, "\t", c.Intelligence.GetModifier())
	fmt.Fprintln(writer, "\tCharisma:\t", c.Charisma.Score, "\t", c.Charisma.GetModifier())
	writer.Flush()

	fmt.Println()
	fmt.Println("Equipment:")
	for _, item := range c.equipment {
		fmt.Printf("\t%+v\n", item)
	}
	fmt.Println()
}
