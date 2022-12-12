package goaocd

// Position with coordinates X and Y
type Pos struct {
	X, Y int
}

// Create new position by applying mutation
func (p *Pos) Mut(m Pos) Pos {
	return Pos{p.X + m.X, p.Y + m.Y}
}

// Check position is equals to another one
func (p *Pos) Eq(m Pos) bool {
	return p.X == m.X && p.Y == m.Y
}

// Calculate Manhattan distance between 2 positions
func (p *Pos) ManhattanDist(to Pos) int {
	return absDiffInt(p.X, to.X) + absDiffInt(p.Y, to.Y)
}
