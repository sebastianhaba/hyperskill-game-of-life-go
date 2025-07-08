package universe

import (
	"fmt"
	"math/rand"
)

type Universe struct {
	Size       int
	seed       int64
	rnd        *rand.Rand
	current    []int
	generation int
}

type Neighbour struct {
	Row    int
	Column int
}

func NewUniverse(size int, seed int64) *Universe {
	u := &Universe{
		Size:       size,
		seed:       seed,
		current:    make([]int, size*size),
		generation: 1,
	}
	u.rnd = rand.New(rand.NewSource(seed))
	for i := 0; i < size*size; i++ {
		u.current[i] = u.rnd.Intn(2)
	}

	return u
}

func (u *Universe) Display() {
	fmt.Println("Generation: #", u.GetGeneration())
	fmt.Println("Alive: ", u.GetAlive())
	for row := 0; row < u.Size; row++ {
		for col := 0; col < u.Size; col++ {
			if u.IsAlive(row, col) {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func (u *Universe) IsAlive(row, col int) bool {
	return u.current[row*u.Size+col] == 1
}

func (u *Universe) GetNeighbours(row, col int) []Neighbour {
	neighbours := make([]Neighbour, 0, 8)
	directions := []struct{ dRow, dCol int }{
		{-1, 0},  // N
		{-1, 1},  // NE
		{0, 1},   // E
		{1, 1},   // SE
		{1, 0},   // S
		{1, -1},  // SW
		{0, -1},  // W
		{-1, -1}, // NW
	}

	for _, dir := range directions {
		nRow := (row + dir.dRow + u.Size) % u.Size
		nCol := (col + dir.dCol + u.Size) % u.Size
		neighbours = append(neighbours, Neighbour{Row: nRow, Column: nCol})
	}

	return neighbours
}

func (u *Universe) GetNumAliveNeighbours(row, col int) int {
	neighbours := u.GetNeighbours(row, col)
	aliveNeighbours := 0
	for _, n := range neighbours {
		if u.IsAlive(n.Row, n.Column) {
			aliveNeighbours++
		}
	}

	return aliveNeighbours
}

func (u *Universe) NextGeneration() {
	next := make([]int, u.Size*u.Size)
	for row := 0; row < u.Size; row++ {
		for col := 0; col < u.Size; col++ {
			aliveNeighbours := u.GetNumAliveNeighbours(row, col)
			if u.IsAlive(row, col) {
				if aliveNeighbours < 2 || aliveNeighbours > 3 {
					next[row*u.Size+col] = 0
				} else {
					next[row*u.Size+col] = 1
				}
			} else {
				if aliveNeighbours == 3 {
					next[row*u.Size+col] = 1
				} else {
					next[row*u.Size+col] = 0
				}
			}
		}
	}

	u.current = next
	u.generation++
}

func (u *Universe) GetAlive() int {
	alive := 0

	for row := 0; row < u.Size; row++ {
		for col := 0; col < u.Size; col++ {
			if u.IsAlive(row, col) {
				alive++
			}
		}
	}

	return alive
}

func (u *Universe) GetGeneration() int {
	return u.generation
}
