package character

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

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

	if hp := int(c.class.GetHitDie()) + c.constitution.GetModifier(); hp < 1 {
		return nil, ErrInvalidHP
	} else {
		c.hp.Max = uint(hp)
	}
	c.hp.Current = int(c.hp.Max)

	c.inventory = make(map[equipment.Equipment]int)

	return c, nil
}

func (c *Character) GetSpeed() uint {
	return c.race.Speed
}

func (c *Character) GetAC() uint {
	ac := 10 + c.dexterity.GetModifier() + c.race.Size.GetACBonus()

	if ac < 0 {
		return 0
	}

	return uint(ac)
}

// class
func (c *Character) LevelUp(takeAverage bool) {
	c.class.LevelUp()
	hpGain := c.class.GetHPRoll()
	if takeAverage {
		hpGain = c.class.GetHPAverage()
	}
	c.hp.ModifyMax(int(hpGain))
}

// saves
func (c *Character) GetFortitudeSave() int {
	return c.constitution.GetModifier()
}
func (c *Character) GetReflexSave() int {
	return c.dexterity.GetModifier()
}
func (c *Character) GetWillSave() int {
	return c.wisdom.GetModifier()
}

// inventory
func (c *Character) GetInventory() map[equipment.Equipment]int {
	return c.inventory
}
func (c *Character) AddItem(i equipment.Equipment) {
	c.inventory[i]++
}
func (c *Character) RemoveItem(i equipment.Equipment) error {
	if v, ok := c.inventory[i]; ok {
		c.inventory[i]--
		if v-1 < 1 {
			delete(c.inventory, i)
		}
		return nil
	}
	return errors.New("Item not in inventory")
}

// HP
// GetHP returns the current and max HP of the character.
func (c *Character) GetHP() stats.HP {
	return c.hp
}

// Damage the character for hp. If hp is negative, return an error.
func (c *Character) Damage(hp uint) {
	c.hp.ModifyCurrent(-int(hp))
}

// Heal the character for hp. If hp is negative, return an error.
func (c *Character) Heal(hp int) error {
	if hp < 0 {
		return ErrInvalidHP
	}
	c.hp.ModifyCurrent(hp)
	return nil
}

// display
func (c *Character) Print() {
	fmt.Printf("Name: %s \t Race: %s\n", c.name, c.race.Name)
	fmt.Printf("Age: %d \t Height: %s\t Weight: %d\n", c.age, c.height, c.weight)
	fmt.Printf("Level: %d \t Class: %s\n", c.class.GetLevel(), c.class.Name)

	fmt.Printf("HP: %d/%d\n", c.hp.Current, c.hp.Max)
	fmt.Printf("AC: %d\n", c.GetAC())
	fmt.Printf("Speed: %d\n", c.GetSpeed())

	fmt.Printf("Stats:\n")
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "\tStrength:\t", c.strength.Score, "\t", c.strength.GetModifier())
	fmt.Fprintln(writer, "\tDexterity:\t", c.dexterity.Score, "\t", c.dexterity.GetModifier())
	fmt.Fprintln(writer, "\tConstitution:\t", c.constitution.Score, "\t", c.constitution.GetModifier())
	fmt.Fprintln(writer, "\tWisdom:\t", c.wisdom.Score, "\t", c.wisdom.GetModifier())
	fmt.Fprintln(writer, "\tIntelligence:\t", c.intelligence.Score, "\t", c.intelligence.GetModifier())
	fmt.Fprintln(writer, "\tCharisma:\t", c.charisma.Score, "\t", c.charisma.GetModifier())
	writer.Flush()

	fmt.Println()
	fmt.Println("Inventory:")
	for item, v := range c.inventory {
		fmt.Printf("\t%+v:\t%d\n", item, v)
	}
	fmt.Println()
	if c.primaryHand != nil {
		fmt.Println("Primary Hand:", c.primaryHand.GetName())
	}
	fmt.Println()
}

// test
func (c *Character) SetPrimaryHand(i equipment.Equipment) error {
	if _, ok := c.inventory[i]; ok {
		c.primaryHand = i
		return nil
	}
	return errors.New("Item not in inventory")
}

func (c *Character) Attack(target *Character) {
	if w, ok := c.primaryHand.(equipment.Weapon); ok {
		target.Damage(w.GetDamage())
	}
}
