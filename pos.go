package goaocd

import "strings"

// Position with coordinates X and Y
type Pos struct {
	X, Y int
}

// Create new Pos from string of format XXX,YYY
// Example:
//
//	pos := NewPos("10,20")
func NewPos(s string) Pos {
	coords := strings.Split(s, ",")
	x := atoi(coords[0])
	y := atoi(coords[1])
	return Pos{X: x, Y: y}
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
