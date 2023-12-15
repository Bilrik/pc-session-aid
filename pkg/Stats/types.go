package stats

type Ability struct {
	Score    uint
	Modifier int
}

type HP struct {
	Current int
	Max     uint
}

type Size struct {
	name  string
	ACmod int
}
