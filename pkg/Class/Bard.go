package class

var Bard = &bard{
	class: class{
		Name:   "Bard",
		hitDie: 8,
		level:  1,
	},
    caster: caster{
        spellsKnown: []uint{},
        spellSlots: []uint{},
        maxSpellSlots: []uint{},
    },
}

type bard struct {
	class
    caster  
}
