package items

import (
    equipment "github.com/CyTechNomad/pc-session-aid/pkg/Equipment"
    "errors"
)

var Backpack = backpack{
    Item: equipment.Item{
        Name: "Backpack",
        Description: "A backpack to hold all your stuff",
        Weight: 5,
        Cost: 2,
    },
    maxCapacity: 10,
    items: []equipment.Equipment{},
}

type backpack struct {
    equipment.Item
    maxCapacity uint
    items    []equipment.Equipment
}

func (b backpack) GetCapacity() uint {
    return b.maxCapacity - uint(len(b.items))
}

func (b backpack) GetMaxCapacity() uint {
    return b.maxCapacity
}

func (b backpack) GetItems() []equipment.Equipment {
    return b.items
}   

func (b *backpack) AddItem(i equipment.Equipment) error {
    if len(b.items) >= int(b.maxCapacity) {
        return errors.New("Backpack is full")
    }
    b.items = append(b.items, i)
    return nil
}

func (b *backpack) RemoveItem(i equipment.Equipment) error {
    for idx, item := range b.items {
        if item == i {
            b.items = append(b.items[:idx], b.items[idx+1:]...)
            return nil
        }
    }
    return errors.New("Item not in backpack")
}

func (b *backpack) GetWeight() float32 {
    weight := b.Weight
    for _, item := range b.items {
        weight += item.GetWeight()
    }
    return weight
}
