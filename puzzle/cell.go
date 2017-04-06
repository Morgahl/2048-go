package puzzle

import (
	"fmt"
)

type cell struct {
	value uint
}

func newCell() *cell {
	return &cell{}
}

func (c cell) String() string {
	if c.isEmpty() {
		return "-"
	}

	return fmt.Sprintf("%d", c.value)
}

func (c cell) isEmpty() bool {
	return c.value == 0
}

// merges cells in left to right order
func mergeCells(cs []*cell) bool {
	var work bool

	// pack initial cells towards left
	for l := 0; l < len(cs); l++ {
		csl := cs[l]
		if csl.isEmpty() {
			for r := l + 1; r < len(cs); r++ {
				csr := cs[r]
				if !csr.isEmpty() {
					csl.value = csr.value
					csr.value = 0
					work = true
					break
				}
			}
		}
	}

	// check for cell merges and merge right into left
	for l := 0; l < len(cs)-1; l++ {
		left := cs[l]
		if left.isEmpty() {
			break
		}

		right := cs[l+1]
		if right.isEmpty() {
			break
		}

		if left.value == right.value {
			left.value *= 2
			right.value = 0
			work = true
			l++
		}
	}

	// pack final cells towards left
	for l := 0; l < len(cs); l++ {
		csl := cs[l]
		if csl.isEmpty() {
			for r := l + 1; r < len(cs); r++ {
				csr := cs[r]
				if !csr.isEmpty() {
					csl.value = csr.value
					csr.value = 0
					work = true
					break
				}
			}
		}
	}

	return work
}
