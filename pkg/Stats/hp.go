package stats

func (hp *HP) ModifyCurrent(amount int) {
	hp.Current += amount
}

func (hp *HP) ModifyMax(amount int) {
	if int(hp.Max)+amount < 0 {
		hp.Max = 0
		return
	}
	// TODO: Modify current HP
}
