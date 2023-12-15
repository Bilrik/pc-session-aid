package main

import (
	"fmt"

	character "github.com/Bilrik/pc-session-aid/pkg/Character"
	race "github.com/Bilrik/pc-session-aid/pkg/Race"
)

func main() {
	c := character.NewCharacter(
		character.WithCharacterName("Bilrik"),
		character.WithRace(race.Gnome),
		character.WithAge(30),
		character.WithHeight("6'2\""),
		character.WithAbilityScores(18, 16, 14, 12, 10, 8),
		character.WithMaxHP(10),
	)

	c.HP.ModifyCurrent(-5)
	c.HP.ModifyCurrent(2)

	fmt.Printf("%+v\n", c)
	fmt.Printf("speed: %d\n", c.GetSpeed())
}
