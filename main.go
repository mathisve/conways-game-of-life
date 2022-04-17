package main

import (
	"fmt"
	"math/rand"
	"time"
)

type world struct {
	Size  int
	Board [][]int
}

func create_world(size int) world {
	a := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		a[i] = make([]int, size+1)
	}
	return world{
		Size:  size,
		Board: a,
	}
}

func (w *world) print_world() {
	for i := 1; i < w.Size; i++ {
		for j := 1; j < w.Size; j++ {
			if w.Board[i][j] == 0 {
				fmt.Print("·")
			} else {
				fmt.Print("✖")
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func (w *world) fill_world() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < w.Size; i++ {
		for j := 1; j < w.Size; j++ {
			if rand.Intn(100) < 10 {
				w.Board[i][j] = 1
			}
		}
	}
}

func (w *world) get_surrounding_variables(x int, y int) (surrounding_a [][]int) {
	surrounding_a = make([][]int, 3)
	for i := 0; i < 3; i++ {
		surrounding_a[i] = w.Board[x-1+i][y-1 : y+2]
	}
	return surrounding_a
}

func sum_surrounding_variables(surrounding_a [][]int) (sum int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += surrounding_a[i][j]
		}
	}
	sum -= surrounding_a[1][1]
	return sum
}

func (w *world) update_world() {
	updated_world := w

	for i := 1; i < updated_world.Size; i++ {
		for j := 1; j < updated_world.Size; j++ {
			sum := sum_surrounding_variables(updated_world.get_surrounding_variables(i, j))

			if updated_world.Board[i][j] == 1 {
				if sum == 2 || sum == 3 {
					updated_world.Board[i][j] = 1
				} else {
					updated_world.Board[i][j] = 0
				}
			} else {
				if sum == 3 {
					updated_world.Board[i][j] = 1
				}
			}
		}
	}

	w = &world{
		Size:  updated_world.Size,
		Board: updated_world.Board,
	}
}

func main() {
	size := 25
	steps := 100

	a := create_world(size)
	a.fill_world()

	a.print_world()
	for i := 0; i < steps; i++ {
		fmt.Println(i)
		a.update_world()
		fmt.Println("\033[H\033[2J")
		a.print_world()
		time.Sleep(200 * time.Millisecond)
	}

	a.print_world()
}
