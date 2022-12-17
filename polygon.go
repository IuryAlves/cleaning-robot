package cleaning_robot

type Coordinate struct {
	X, Y int
}

type Polygon struct {
	Coordinates []Coordinate
}

// Area calculates the area using the Gaussian method
func (p *Polygon) Area() int {
	if len(p.Coordinates) == 0 {
		return 0
	}
	var xSum, ySum int
	c := p.Coordinates
	c = append(c, p.Coordinates[0])
	for i := 0; i < len(c)-1; i++ {
		xSum += c[i].X * c[i+1].Y
		ySum += c[i].Y * c[i+1].X
	}
	return (xSum - ySum) * -1
}

func (p *Polygon) Add(x, y int) {
	p.Coordinates = append(
		p.Coordinates,
		Coordinate{
			X: x,
			Y: y,
		},
	)
}
func (p *Polygon) Length() int {
	return len(p.Coordinates)
}
