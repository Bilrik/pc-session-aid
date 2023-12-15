package stats

func (a *Ability) SetAbilityScore(score uint) {
	a.Score = score
	a.Modifier = int(score/2 - 5)
}
