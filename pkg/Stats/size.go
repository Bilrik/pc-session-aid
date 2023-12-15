package stats

var (
	Small = Size{
		name:  "Small",
		ACmod: 1,
	}
	Medium = Size{
		name:  "Medium",
		ACmod: 0,
	}
	Large = Size{
		name:  "Large",
		ACmod: -1,
	}
)

func (s *Size) GetACBonus() int {
	return s.ACmod
}
