package stats

func (a *Ability) GetModifier() int {
	return a.Modifier
}

func (a *Ability) SetAbilityScore(score uint) {
	a.Score = score
	a.Modifier = int(score/2 - 5)
}
