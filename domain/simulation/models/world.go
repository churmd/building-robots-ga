package models

type Coord struct {
	X int
	Y int
	Z int
}

type Block struct {
	Location Coord
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Robot struct {
	Location Coord
	Direction
}

type World struct {
	// The coordinate trying to be reached
	Goal Coord
	// The upper right of the world, lower left of the world being {0,0,0}
	MaxCoord Coord

	Blocks []Block

	Robots []Robot
}
