package magic

type Castable interface{
    GetName() string
    GetDescription() string
    GetLevel() uint
}

type ArcaneSpell interface{
    GetSchool() string
}

type DivineSpell interface{
    GetDomain() string
}

type DamageSpell interface{
    GetDamage() uint
}

type RangeSpell interface{
    GetRange() uint
}

type Spell struct{
    Name string
    Description string
    Level uint
}

func (s Spell) GetName() string{
    return s.Name
}

func (s Spell) GetDescription() string{
    return s.Description
}

func (s Spell) GetLevel() uint{
    return s.Level
}
