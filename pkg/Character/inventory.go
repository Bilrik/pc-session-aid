package character

import (
	"errors"

	equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"
)

func (c *Character) GetInventory() map[equipment.Equipment]int {
	return c.inventory
}

func (c *Character) AddItem(i equipment.Equipment) {
	c.inventory[i]++
}

func (c *Character) AddItems(i equipment.Equipment, count int) {
	c.inventory[i] += count
}

func (c *Character) RemoveItem(i equipment.Equipment) error {
	if v, ok := c.inventory[i]; ok {
		c.inventory[i]--
		if v-1 < 1 {
			delete(c.inventory, i)
		}
		return nil
	}
	return errors.New("Item not in inventory")
}

func (c *Character) RemoveItems(i equipment.Equipment, count int) error {
	if v, ok := c.inventory[i]; ok {
		c.inventory[i] -= count
		if v-count < 1 {
			delete(c.inventory, i)
		}
		return nil
	}
	return errors.New("Item not in inventory")
}

func (c *Character) RemoveAllItem(i equipment.Equipment) error {
	if _, ok := c.inventory[i]; ok {
		delete(c.inventory, i)
		return nil
	}
	return errors.New("Item not in inventory")
}

func (c *Character) ClearInventory() {
	c.inventory = make(map[equipment.Equipment]int)
}
