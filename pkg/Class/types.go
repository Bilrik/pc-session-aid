package class

import magic "github.com/CyTechNomad/pc-session-aid/pkg/Magic"

type Class interface {
	GetLevel() uint
	LevelUp()
	GetHitDie() uint
	GetName() string
}

type class struct {
	Name   string
	hitDie uint
	level  uint
}

func (c *class) GetName() string {
    return c.Name
}

func (c *class) GetLevel() uint {
    return c.level
}

func (c *class) LevelUp() {
    c.level++
}

func (c *class) GetHitDie() uint {
    return c.hitDie
}

type Caster interface {
    getSpellSlots() []uint
    getSpellsKnown() [][]uint
}

type caster struct {
    spellsKnown []uint
    spellSlots []uint
    maxSpellSlots []uint
}

func (c *caster) getSpellSlots() []uint {
    return c.spellSlots
}

func (c *caster) getSpellsKnown() []uint {
    return c.spellsKnown
}

func (c *caster) consumeSpellSlot(level uint) error{
    if c.spellSlots[level] > 0 {
        c.spellSlots[level]--
        return nil
    }
    return ErrNoSpellSlots    
}

func (c *caster) regainSpellSlot(level uint) error{
    if c.spellSlots[level] < c.maxSpellSlots[level] {
        c.spellSlots[level]++
        return nil
    }
    return ErrMaxSpellSlots
}

type PreparedCaster interface {
    GetPreparedSpells() [][]magic.Spell
    PrepareSpell(spell magic.Spell) error
    consumePreparedSpell(spell magic.Spell) error
}
