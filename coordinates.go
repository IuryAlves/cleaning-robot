package cleaning_robot

type Coordinate struct {
	X, Y int
}

type Coordinates struct {
	coordinates []Coordinate
}

func (p *Coordinates) Add(x, y int) {
	p.coordinates = append(
		p.coordinates,
		Coordinate{
			X: x,
			Y: y,
		},
	)
}
func (p *Coordinates) Length() int {
	return len(p.coordinates)
}
