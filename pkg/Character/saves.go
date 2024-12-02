package character

// saves
func (c *Character) GetFortitudeSave() int {
	return c.constitution.GetModifier()
}
func (c *Character) GetReflexSave() int {
	return c.dexterity.GetModifier()
}
func (c *Character) GetWillSave() int {
	return c.wisdom.GetModifier()
}
