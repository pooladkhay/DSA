package main

import (
	"errors"
	"fmt"
)

type coordinate struct {
	x int
	y int
}

// Represents a point in the grid, either '.' or '%'
type Point struct {
	coord       coordinate
	visited     bool
	hasAsteroid bool
}

type Graph struct {
	// 2D grid representing the input
	grid [][]*Point
	// List of points that are potentially part of an asteroid
	asteroids []*Point
}

// Builds a Graph from a given input.
// All lines must be of the same length.
// Example input:
//
//	input := []string{
//		"................",
//		"....%.......%...",
//		"....%%..........",
//		".....%.....%....",
//		"....%.......%...",
//		"...%.%..........",
//		"......%........%",
//	}
func BuildGraph(input []string) (*Graph, error) {
	colLen := len(input)
	rowLen := len(input[0])

	grid := make([][]*Point, 0, colLen)
	asteroids := make([]*Point, 0, colLen)

	for y, line := range input {

		if len(line) != rowLen {
			return nil, fmt.Errorf("error parsing index number %d in the input slice. all entries in the input slice must be of the same length", y)
		}

		row := make([]*Point, 0, rowLen)
		for x, p := range line {

			point := Point{
				coord:       coordinate{x, y},
				hasAsteroid: false,
				visited:     false,
			}
			if p == '%' {
				point.hasAsteroid = true
				asteroids = append(asteroids, &point)
			}

			row = append(row, &point)
		}

		grid = append(grid, row)
	}

	return &Graph{grid, asteroids}, nil
}

// Returns a pointer to a Point with coordinate (c.x, c.y), if there is an asteroids on that point.
func (g *Graph) isAsteroidAt(c coordinate) (*Point, error) {
	if g.grid[c.y][c.x].hasAsteroid {
		return g.grid[c.y][c.x], nil
	}
	return nil, errors.New("no asteroid on this coordinate")
}

// Returns a list of nearby points to p that contain an asteroid and are not visited yet.
func (g *Graph) GetNearbyAsteroids(p *Point) ([]*Point, error) {
	res := []*Point{}

	/*
		// 	                 |
		// 	              (0, 1)
		// 	    (-1, 1)*     *     *(1, 1)
		// 	                 |
		//  ----(-1, 0)*-----+-----*(1, 0)-----
		// 	                 |
		//      (-1,-1)*     *     *(1, -1)
		// 	              (0, -1)
		// 	                 |
		// 	                 |
	*/
	possibleTransformations := []struct {
		x int
		y int
	}{
		{x: 1, y: 0},
		{x: 1, y: 1},
		{x: 0, y: 1},
		{x: -1, y: 1},
		{x: -1, y: 0},
		{x: -1, y: -1},
		{x: 0, y: -1},
		{x: 1, y: -1},
	}

	for _, t := range possibleTransformations {
		coo := coordinate{x: p.coord.x + t.x, y: p.coord.y + t.y}
		// Ignore a coordinate if it has a negative value on either x or y
		// or is outside of the plane coordinates after the transformation
		if coo.x >= 0 && coo.y >= 0 && coo.x < len(g.grid[0]) && coo.y < len(g.grid) {
			if p, err := g.isAsteroidAt(coo); err == nil {
				if !p.visited {
					res = append(res, p)
				}
			}
		}
	}

	if len(res) > 0 {
		return res, nil
	}

	return nil, errors.New("no asteroids nearby")
}

// Prints the Graph with asteroids.
func (g *Graph) Print() {
	for _, row := range g.grid {
		for _, point := range row {
			if point.hasAsteroid {
				fmt.Printf("[%v]", "%")
			} else {
				fmt.Printf(" %v ", ".")
			}
		}
		fmt.Printf("\n")
	}
}
