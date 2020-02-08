package main

import (
	"fmt"
	"math/rand"
	"time"
)

var size int
var fill_odds int

func create_world () ([][]int) {
	a := make([][]int, size+1)
	for i:=0; i<size+1; i++ {
		a[i] = make([]int, size+1)
	}
	return a
}


func print_world (a [][]int) {
	for i:=1; i<size; i++ {
		for j:=1; j<size; j++ {
			if a[i][j] == 0 {
				fmt.Print("●")
			} else {
				fmt.Print("✖")
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func fill_world (a [][]int) ([][]int) {
	rand.Seed(time.Now().UnixNano())
	for i :=1; i<size; i++ {
		for j:=1; j<size; j++ {
			if rand.Intn(100) < fill_odds {
				a[i][j] = 1
			}
		}
	}
	return a
}


func main () {
	size = 20
	fill_odds = 25
	a := create_world()
	a = fill_world(a)
	print_world(a)
	

}
