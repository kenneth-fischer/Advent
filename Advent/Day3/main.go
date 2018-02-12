package main

import (
	"fmt"
	"math"
)

type memmap map[int]cell
type locations map[int]map[int]int

type vector2D struct {
	X int
	Y int
}

func (v vector2D) Add(v2 vector2D) vector2D {
	return vector2D{v.X + v2.X, v.Y + v2.Y}
}

func (v vector2D) ToOrigin() int {
	return int(math.Abs(float64(v.X)) + math.Abs(float64(v.Y)))
}

func (v vector2D) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

type cell struct {
	Location vector2D
	NextDir  vector2D
	Value    int
}

var (
	right   = vector2D{1, 0}
	up      = vector2D{0, 1}
	left    = vector2D{-1, 0}
	down    = vector2D{0, -1}
	mem     = memmap{}
	locs    = locations{}
	currDir vector2D
)

func main() {
	puzzleInput := 277678
	c := mapAddr(puzzleInput)

	// Print graphical depcition of addresses
	for row := 2; row >= -2; row-- {
		for col := -2; col <= 2; col++ {
			addr := locs[col][row]
			fmt.Printf(" %02d", addr)
		}
		fmt.Println()
	}
	// Print distance from origin for the address equal to the given puzzle input
	fmt.Printf("Distance(%d) = %d\n", puzzleInput, c.Location.ToOrigin())

	// Print graphical depiction of cell values
	for row := 2; row >= -2; row-- {
		for col := -2; col <= 2; col++ {
			addr := locs[col][row]
			cell := mem[addr]
			fmt.Printf(" %3d", cell.Value)
		}
		fmt.Println()
	}

	// Print first value larger than puzzle input value
	for i := 1; i > 0; i++ {
		if mapAddr(i).Value > puzzleInput {
			fmt.Printf("%d is first value larger than %d\n", mapAddr(i).Value, puzzleInput)
			break
		}
	}
}

func mapAddr(addr int) cell {
	if existing, found := mem[addr]; found {
		return existing
	}

	var curr cell

	if addr == 1 {
		curr = cell{vector2D{0, 0}, right, 1}
	} else {
		prev := mapAddr(addr - 1)
		curr = nextCell(prev)
		curr.Value = valueOf(curr)
	}
	associate(addr, curr)
	return curr
}

func associate(addr int, c cell) {
	mem[addr] = c
	if _, exists := locs[c.Location.X]; !exists {
		locs[c.Location.X] = map[int]int{}
	}
	locs[c.Location.X][c.Location.Y] = addr
}

func locationAllocated(v vector2D) bool {
	row, exists := locs[v.X]
	if !exists {
		return false
	}
	if _, exists = row[v.Y]; !exists {
		return false
	}
	return true
}

func inLocation(loc vector2D) *cell {
	row, exists := locs[loc.X]
	if !exists {
		return nil
	}
	addr, exists := row[loc.Y]
	if !exists {
		return nil
	}
	c := mem[addr]
	return &c

}

func valueOf(c cell) int {
	value := 0

	for _, nbr := range neighbors(c) {
		value += nbr.Value
	}
	return value
}

func neighbors(c cell) []cell {
	cells := []cell{}

	if nbr := inLocation(c.Location.Add(up).Add(left)); nbr != nil {
		cells = append(cells, *nbr)
	}
	if nbr := inLocation(c.Location.Add(up)); nbr != nil {
		cells = append(cells, *nbr)
	}
	if nbr := inLocation(c.Location.Add(up).Add(right)); nbr != nil {
		cells = append(cells, *nbr)
	}

	if nbr := inLocation(c.Location.Add(left)); nbr != nil {
		cells = append(cells, *nbr)
	}
	if nbr := inLocation(c.Location.Add(right)); nbr != nil {
		cells = append(cells, *nbr)
	}

	if nbr := inLocation(c.Location.Add(down).Add(left)); nbr != nil {
		cells = append(cells, *nbr)
	}
	if nbr := inLocation(c.Location.Add(down)); nbr != nil {
		cells = append(cells, *nbr)
	}
	if nbr := inLocation(c.Location.Add(down).Add(right)); nbr != nil {
		cells = append(cells, *nbr)
	}

	return cells
}

func nextCell(prev cell) cell {
	currLoc := prev.Location.Add(prev.NextDir)
	// Try to continue to "curl" the spiral by changing the direction.
	dir := nextDir(prev.NextDir)

	nextLoc := currLoc.Add(dir)
	// If we are trying to curl the spiral back on itself then restore the current direction
	if locationAllocated(nextLoc) {
		dir = prev.NextDir
	}

	return cell{currLoc, dir, 0}
}

func nextDir(currDir vector2D) vector2D {
	switch currDir {
	case right:
		return up
	case up:
		return left
	case left:
		return down
	default:
		return right
	}
}
