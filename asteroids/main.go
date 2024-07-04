package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	input := []string{
		"................",
		".....%%.........",
		"...%%.......%%..",
		"....%%........%.",
		".....%.....%....",
		"....%.......%...",
		"...%.%.%........",
		"...%..%.%......%",
		"................",
	}

	field, err := BuildGraph(input)
	if err != nil {
		log.Fatalln(err)
	}

	asteroidLengthList := []int{}

	for _, asteroid := range field.asteroids {
		// Slightly modified BFS on each '%' point, if not visited yet.
		// The modification is that instead of marking each point as visited
		// after dequeue operation, the algorithm marks them as visited right
		// before enqueue operation.
		if !asteroid.visited {
			// Represents total number of points connected together to form an asteroid
			asteroidLength := 0

			queue := make([]*Point, 0, len(field.asteroids))
			queue = append(queue, asteroid)
			asteroid.visited = true

			for len(queue) > 0 {
				// Dequeue
				ast := queue[0]
				queue = queue[1:]

				asteroidLength++

				if nearbyAsteroids, err := field.GetNearbyAsteroids(ast); err == nil {
					for _, a := range nearbyAsteroids {
						a.visited = true
						// Enqueue
						queue = append(queue, a)
					}
				}
			}

			asteroidLengthList = append(asteroidLengthList, asteroidLength)
		}
	}

	// Deal with the the result
	if len(asteroidLengthList) > 0 {
		field.Print()

		sort.Ints(asteroidLengthList)

		fmt.Printf("Found asteroids of length: %v\n", asteroidLengthList)
		fmt.Println("Largest:", asteroidLengthList[len(asteroidLengthList)-1])
	} else {
		fmt.Println("No asteroids were found.")
	}
}
