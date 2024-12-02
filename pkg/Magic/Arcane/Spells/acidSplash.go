package spells

import (
    "github.com/CyTechNomad/pc-session-aid/pkg/Magic"
)

var AcidSplash = acidSplash{
    Spell: magic.Spell{ 
        Name: "Acid Splash",   
        Description: "You hurl a bubble of acid. You must succeed on a ranged touch attack to hit your target. The acid deals 1d3 points of acid damage.",
        Level: 0,
    },
}

type acidSplash struct {
    magic.Spell
    school string
}

func (a acidSplash) GetSchool() string {
    return "Conjuration"
}

func (a acidSplash) GetDamage() uint {
    return 1
}

func (a acidSplash) GetRange() uint {
    return 25
}
