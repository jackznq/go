package znq

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a Artist) int {
	a.Songs++
	return a.Songs
}
