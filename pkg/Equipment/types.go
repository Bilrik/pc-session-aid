package equipment

type Equipment interface {
	GetName() string
	GetDescription() string
	GetWeight() float32
	GetCost() float32
}

type Weapon interface {
	GetDamage() uint
}

type RangedWeapon interface {
	RequiredAmmo() string
}

type ChargeWeapon interface {
	GetCharges() uint
	UseCharge()
}

type AcItem interface {
	GetACBonus() int
}

type Ammunition interface {
	GetAmmoType() string
}

type Container interface {
    GetCapacity() uint
    GetItems() []Equipment
    AddItem(Equipment) error
    RemoveItem(Equipment) error
    GetWeight() float32
}

type Item struct {
	Name        string
	Description string
	Weight      float32
	Cost        float32
}

func (i Item) GetName() string {
	return i.Name
}

func (i Item) GetDescription() string {
	return i.Description
}

func (i Item) GetWeight() float32 {
	return i.Weight
}

func (i Item) GetCost() float32 {
	return i.Cost
}
