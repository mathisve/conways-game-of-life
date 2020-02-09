package main

import (
	"fmt"
	"math/rand"
	"time"
)

var size int
var fill_odds int

func create_world() [][]int {
	a := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		a[i] = make([]int, size+1)
	}
	return a
}

func print_world(a [][]int) {
	for i := 1; i < size; i++ {
		for j := 1; j < size; j++ {
			if a[i][j] == 0 {
				fmt.Print("·")
			} else {
				fmt.Print("✖")
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func fill_world(a [][]int) [][]int {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < size; i++ {
		for j := 1; j < size; j++ {
			if rand.Intn(100) < fill_odds {
				a[i][j] = 1
			}
		}
	}
	return a
}

func get_surrounding_variables(x int, y int, a [][]int) (surrounding_a [][]int) {
	surrounding_a = make([][]int, 3)
	for i := 0; i < 3; i++ {
		surrounding_a[i] = a[x-1+i][y-1 : y+2]
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

func update_world(a [][]int) [][]int {
	updated_a := a

	for i := 1; i < size; i++ {
		for j := 1; j < size; j++ {
			sum := sum_surrounding_variables(get_surrounding_variables(i, j, a))

			if a[i][j] == 1 {
				if sum == 2 || sum == 3 {
					updated_a[i][j] = 1
				} else {
					updated_a[i][j] = 0
				}
			} else {
				if sum == 3 {
					updated_a[i][j] = 1
				}
			}
		}
	}
	return updated_a
}

func main() {
	size = 45
	fill_odds = 10
	steps := 100

	a := create_world()
	a = fill_world(a)

	print_world(a)
	for i := 0; i < steps; i++ {
		fmt.Println(i)
		a = update_world(a)
		fmt.Println("\033[H\033[2J")
		print_world(a)
		time.Sleep(1 * time.Second)
	}
}
