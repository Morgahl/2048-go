package puzzle

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	GameExit     = fmt.Errorf("Game exit!")
	GameOver     = fmt.Errorf("Game over!")
	InvalidInput = fmt.Errorf("Invalid Input!")
	Victory      = fmt.Errorf("You Win!")
)

type Puzzle struct {
	victory uint
	sizeX   uint
	sizeY   uint

	cells [][]*cell

	rand *rand.Rand
}

func New(sizeX, sizeY, victory uint, seed int64) *Puzzle {
	p := &Puzzle{
		victory: victory,
		sizeX:   sizeX,
		sizeY:   sizeY,
		cells:   make([][]*cell, sizeX),
		rand:    rand.New(rand.NewSource(seed)),
	}

	for x := uint(0); x < p.sizeX; x++ {
		col := make([]*cell, 0, p.sizeY)
		for y := uint(0); y < p.sizeY; y++ {
			col = append(col, newCell())
		}
		p.cells[x] = col
	}

	p.populateCells(2)

	return p
}

func (p Puzzle) String() string {
	rows := make([]string, p.sizeY)

	for y := uint(0); y < p.sizeY; y++ {
		col := make([]string, p.sizeX)
		for x := uint(0); x < p.sizeX; x++ {
			col[x] = p.cells[x][y].String()
		}
		rows[y] = "\t" + strings.Join(col, "\t")
	}

	return strings.Join(rows, "\n")
}

func (p Puzzle) isSolved() bool {
	for x := uint(0); x < p.sizeX; x++ {
		for y := uint(0); y < p.sizeY; y++ {
			if p.cells[x][y].value == p.victory {
				return true
			}
		}
	}

	return false
}

func (p Puzzle) getEmptyCells() []*cell {
	empty := make([]*cell, 0, p.sizeX*p.sizeY)
	for x := uint(0); x < p.sizeX; x++ {
		col := p.cells[x]
		for y := uint(0); y < p.sizeY; y++ {
			cell := col[y]
			if cell.isEmpty() {
				empty = append(empty, cell)
			}
		}
	}

	return empty
}

func (p *Puzzle) populateCells(count uint) error {
	cells := p.getEmptyCells()

	for i := uint(0); i < count; i++ {
		if len(cells) <= 0 {
			return GameOver
		}

		idx := p.rand.Intn(len(cells))
		cells[idx].value = uint((p.rand.Intn(2) + 1) * 2)

		// delete populated cell from list, in this case we really kinda don't care about order
		cells[idx] = cells[len(cells)-1]
		cells[len(cells)-1] = nil // for GC reasons
		cells = cells[:len(cells)-1]
	}

	return nil
}

func (p *Puzzle) Shift(input string) error {
	var err error
	switch input {
	default:
		err = InvalidInput

	case "\n":
		err = nil

	case "w", "W":
		err = p.shiftUp()

	case "a", "A":
		err = p.shiftLeft()

	case "s", "S":
		err = p.shiftDown()

	case "d", "D":
		err = p.shiftRight()

	case "x", "X":
		err = GameOver
	}

	if err != nil {
		return err
	}

	if p.isSolved() {
		return Victory
	}

	return nil
}

func (p *Puzzle) shiftUp() error {
	var didWork bool
	for x := uint(0); x < p.sizeX; x++ {
		cells := p.cells[x]
		if mergeCells(cells) && !didWork {
			didWork = true
		}
	}

	if didWork {
		return p.populateCells(1)
	}

	return nil
}

func (p *Puzzle) shiftLeft() error {
	var didWork bool
	for y := uint(0); y < p.sizeY; y++ {
		cells := make([]*cell, 0, p.sizeX)

		for x := 0; x < int(p.sizeX); x++ {
			cells = append(cells, p.cells[x][y])
		}

		if mergeCells(cells) && !didWork {
			didWork = true
		}
	}

	if didWork {
		return p.populateCells(1)
	}

	return nil
}

func (p *Puzzle) shiftDown() error {
	var didWork bool
	for x := uint(0); x < p.sizeX; x++ {
		cells := make([]*cell, 0, p.sizeY)

		for y := int(p.sizeY) - 1; y > -1; y-- {
			cells = append(cells, p.cells[x][y])
		}

		if mergeCells(cells) && !didWork {
			didWork = true
		}
	}

	if didWork {
		return p.populateCells(1)
	}

	return nil
}

func (p *Puzzle) shiftRight() error {
	var didWork bool
	for y := uint(0); y < p.sizeY; y++ {
		cells := make([]*cell, 0, p.sizeX)

		for x := int(p.sizeX) - 1; x > -1; x-- {
			cells = append(cells, p.cells[x][y])
		}

		if mergeCells(cells) && !didWork {
			didWork = true
		}
	}

	if didWork {
		return p.populateCells(1)
	}

	return nil
}
