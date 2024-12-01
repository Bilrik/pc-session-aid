package main

import (
	character "github.com/CyTechNomad/pc-session-aid/pkg/Character"
	class "github.com/CyTechNomad/pc-session-aid/pkg/Class"
	Items "github.com/CyTechNomad/pc-session-aid/pkg/Equipment/Items"
	race "github.com/CyTechNomad/pc-session-aid/pkg/Race"
)

func main() {
	c, _ := character.NewCharacter(
		character.WithCharacterName("Bilrik"),
		character.WithRace(race.Gnome),
		character.WithClass(class.Bard),
		character.WithAge(30),
		character.WithHeight("6'2\""),
		character.WithAbilityScores(18, 16, 14, 12, 10, 8),
	)

	c.AddItem(Items.Luckstone)
	c.AddItem(Items.Torch)
	c.AddItem(Items.Torch)
	c.AddItem(Items.Torch)

	if err := c.SetPrimaryHand(Items.Luckstone); err != nil {
		panic(err)
	}

	b, _ := character.NewCharacter(
		character.WithCharacterName("badguy"),
		character.WithRace(race.Gnome),
		character.WithClass(class.Bard),
		character.WithAge(30),
		character.WithHeight("6'2\""),
		character.WithAbilityScores(18, 16, 14, 12, 10, 8),
	)

	c.Attack(b)

	c.Print()
	b.Print()
}
