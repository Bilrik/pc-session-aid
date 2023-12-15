package main

import (
	"fmt"

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
	fmt.Printf("%+v\n", c)

	c.HP.ModifyCurrent(-5)
	c.HP.ModifyCurrent(2)
	c.LevelUp(true)

	fmt.Printf("speed: %d\n", c.GetSpeed())
	fmt.Printf("AC: %d\n", c.GetAC())
	fmt.Printf("Class: %s, Level: %d\n", c.Class.Name, c.Class.GetLevel())
	fmt.Printf("Fort: %d, Ref: %d, Will: %d\n", c.GetFortitudeSave(), c.GetReflexSave(), c.GetWillSave())
	fmt.Printf("%+v\n", c)
	c.AddItem(equipment.ShortSword)
	c.Print()
}
