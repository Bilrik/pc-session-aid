package main

import (
	character "github.com/Bilrik/pc-session-aid/pkg/Character"
	class "github.com/Bilrik/pc-session-aid/pkg/Class"
	equipment "github.com/Bilrik/pc-session-aid/pkg/Equipment"
	race "github.com/Bilrik/pc-session-aid/pkg/Race"
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
	c.Print()

	c.Damage(5)
	c.Heal(2)
	c.LevelUp(true)
	c.AddItem(equipment.ShortSword)

	c.Print()
}
