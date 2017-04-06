package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	p "github.com/curlymon/2048-go/puzzle"
)

func main() {
	puz := p.New(4, 4, 2048, time.Now().UnixNano())

	var count = 0
	defer func() {
		fmt.Printf("iterations: %d\n", count)
	}()
	showAction(puz)

	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanRunes)
	for scan.Scan() {
		count++
		spacer()

		in := scan.Text()
		if in == "\n" {
			continue
		}

		err := puz.Shift(in)

		switch err {
		default:
			panic(err)

		case p.GameOver:
			showEnd(puz)
			return

		case p.GameExit:
			showExit()
			return

		case p.Victory:
			showVictory(puz)
			return

		case p.InvalidInput:
			showInvalid(puz)

		case nil:
			showAction(puz)
		}
	}
}

func spacer() {
	fmt.Printf("\n\n")
}

func showAction(puz *p.Puzzle) {
	fmt.Println(puz)
	fmt.Println("\nactions: (w, a, s, d, x)?")
}

func showInvalid(puz *p.Puzzle) {
	fmt.Println(puz)
	fmt.Println("\nINVALID INPUT!\nactions: (w, a, s, d, x)?")
}

func showVictory(puz *p.Puzzle) {
	fmt.Println(puz)
	fmt.Println("\nYOU WIN!!")
}

func showEnd(puz *p.Puzzle) {
	fmt.Println(puz)
	fmt.Println("\ngame over!!")
}

func showExit() {
	fmt.Println("\nexited")
}
