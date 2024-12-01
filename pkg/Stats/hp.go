package stats

import "fmt"

func (hp *HP) ModifyCurrent(amount int) {
	fmt.Println("Modifying HP", hp.Current, amount)
	hp.Current += amount
}

func (hp *HP) ModifyMax(amount int) {
	if int(hp.Max)+amount < 0 {
		hp.Max = 0
		return
	}

	hp.Max = uint(int(hp.Max) + amount)
}
