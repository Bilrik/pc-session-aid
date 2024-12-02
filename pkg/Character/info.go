package character

import (
	"fmt"
	"os"
	"text/tabwriter"

	stats "github.com/CyTechNomad/pc-session-aid/pkg/Stats"
)

// GetHP returns the current and max HP of the character.
func (c *Character) GetHP() stats.HP {
	return c.hp
}

func (c *Character) GetSpeed() uint {
	return c.race.GetSpeed()
}

func (c *Character) GetAC() uint {
	ac := 10 + c.dexterity.GetModifier() + c.race.GetSize().GetACBonus()

	if ac < 0 {
		return 0
	}

	return uint(ac)
}

// display
func (c *Character) Print() {
	fmt.Printf("Name: %s \t Race: %s\n", c.name, c.race.GetName())
	fmt.Printf("Age: %d \t Height: %s\t Weight: %d\n", c.age, c.height, c.weight)
	fmt.Printf("Level: %d \t Class: %s\n", c.class.GetLevel(), c.class.GetName())

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
