package equipment

type Equipment struct {
	Name        string
	Description string
	Weight      float32
	Cost        float32
}

type Weapon struct {
	Equipment
	DamageDice         uint
	Critical           uint
	CriticalMultiplier uint
	Type               string
	Range              uint
}

type Armor struct {
	Equipment
	ACBonus      uint
	MaxDex       uint
	CheckPenalty int
	Speed        uint
	Special      string
}
