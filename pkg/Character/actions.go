package character

import (
	"errors"

	equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"
)

// class
func (c *Character) LevelUp(takeAverage bool) {
	c.class.LevelUp()
	hpGain := int(c.class.GetHitDie()) + c.constitution.GetModifier()
	// hpGain := c.class.GetHPRoll()
	if takeAverage {
		// hpGain = c.class.GetHPAverage()
	}
	c.hp.ModifyMax(int(hpGain))
}

// SetPrimaryHand sets the piece of equipment as active in the primary hand
// returns an error if the item is not in the inventory
//
// example:
//
//		c := NewCharacter()
//	    c.AddItem(Items.Torch)
//		if err := c.SetPrimaryHand(Items.Torch); err != nil {
//          panic(err)
//      }
//
func (c *Character) SetPrimaryHand(i equipment.Equipment) error {
	if _, ok := c.inventory[i]; ok {
		c.primaryHand = i
		return nil
	}
	return errors.New("Item not in inventory")
}

// Attack attempts to attack the target character
// returns an error if the character has no weapon equipped
//
// example:
//
//      c := NewCharacter()
//      target := NewCharacter()
//
//      c.AddItem(Items.Sword)
//      if err := c.SetPrimaryHand(Items.Sword); err != nil {
//          panic(err)
//      }
//      if err := c.Attack(target); err != nil {
//          panic(err)
//      }
//
func (c *Character) Attack(target *Character) error {
	w, ok := c.primaryHand.(equipment.Weapon)
	if !ok {
		return errors.New("No weapon equipped")
	}

	// handle ranged weapon
	if w, ok := c.primaryHand.(equipment.RangedWeapon); ok {
		hasRequiredAmmo := false
		// handle ammo check
		for item := range c.inventory {
			if a, ok := item.(equipment.Ammunition); ok {
				if a.GetAmmoType() == w.RequiredAmmo() {
					if err := c.RemoveItem(item); err != nil {
						return err
					}
					hasRequiredAmmo = true
					break
				}
			}
		}

		if !hasRequiredAmmo {
			return errors.New("No ammo")
		}
	}

	// handle charge weapon
	if w, ok := c.primaryHand.(equipment.ChargeWeapon); ok {
		if w.GetCharges() < 1 {
			return errors.New("No charges")
		}
		w.UseCharge()
	}

	target.Damage(w.GetDamage())
	return nil
}
