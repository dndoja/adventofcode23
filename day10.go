package main

import (
	"bufio"
	"fmt"
)

type Location struct {
	pipe  byte
	x     int
	y     int
	north *Location
	west  *Location
	east  *Location
	south *Location
}

type PipesMap = [MapSize][MapSize]byte

const MapSize = 140

func RunDay10(scanner *bufio.Scanner) {
    var pipesMap PipesMap
    var animalLocation Location 

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < MapSize; x++ {
			pipesMap[y][x] = line[x]

            if line[x] == 'S' {
                animalLocation = Location{pipe: 'S', x: x, y: y}
            }
		}
		y++
	}

    furthestDistance := getFurthestDistanceFromS(&animalLocation, pipesMap)
	fmt.Println(furthestDistance)
}

func getFurthestDistanceFromS(startLocation *Location, pipesMap PipesMap) int {
    frontier := []*Location {startLocation}
    distances := map[uint16]int {getLocationId(startLocation): 0}
    furthestDistance := 0

    for len(frontier) > 0 {
        current := frontier[0]
        frontier = frontier[1:]
        neighbours := findNeighbours(current, pipesMap)
        printLocation(current)

        for _, next := range neighbours {
            nextId := getLocationId(next)

            if _, exists := distances[nextId]; !exists {
                frontier = append(frontier, next)
                nextDistance := distances[getLocationId(current)] + 1
                distances[nextId] = nextDistance
                
                if nextDistance > furthestDistance {
                    furthestDistance = nextDistance
                }
            }
        }
    }

   return furthestDistance 
}

func getLocationId(location *Location) uint16 {
    var id uint16
    id = uint16(location.x) << 8
    id = id | uint16(location.y)

    return id
}

func printLocation(location *Location) {
    fmt.Print(" ")
    if location.north != nil {
        fmt.Print(string(location.north.pipe))
    }else{
        fmt.Print(" ")
    }
    fmt.Print(" ")

    fmt.Println("")

    if location.west != nil {
        fmt.Print(string(location.west.pipe))
    }else{
        fmt.Print(" ")
    }
    fmt.Print(string(location.pipe))
    if location.east != nil {
        fmt.Print(string(location.east.pipe))
    }else{
        fmt.Print(" ")
    }

    fmt.Println("")

    fmt.Print(" ")
    if location.south != nil {
        fmt.Print(string(location.south.pipe))
    }else{
        fmt.Print(" ")
    }
    fmt.Println(" ")
    fmt.Println("====")
}

func findNeighbours(location *Location, pipesMap PipesMap) []*Location {
	x := location.x
	y := location.y
    neighbours := []*Location {}
	north := locationAtCoordinates(pipesMap, x, y-1)
	east := locationAtCoordinates(pipesMap, x+1, y)
	south := locationAtCoordinates(pipesMap, x, y+1)
	west := locationAtCoordinates(pipesMap, x-1, y)
    
    if north != nil && fitsNorth(location.pipe) && fitsSouth(north.pipe) {
        location.north = north
        neighbours = append(neighbours, north)
    }

    if south != nil && fitsSouth(location.pipe) && fitsNorth(south.pipe) {
       location.south = south 
       neighbours = append(neighbours, south)
    }

    if west != nil && fitsWest(location.pipe) && fitsEast(west.pipe) {
        location.west = west
        neighbours = append(neighbours, west)
    }

    if east != nil && fitsEast(location.pipe) && fitsWest(east.pipe) {
        location.east = east
        neighbours = append(neighbours, east)
    }

    return neighbours
}

func fitsWest(pipe byte) bool {
   return pipe == 'S' || pipe == '7' || pipe == '-' || pipe == 'J' 
}

func fitsEast(pipe byte) bool {
    return pipe == 'S' || pipe == '-' || pipe == 'F' || pipe == 'L'
}

func fitsNorth(pipe byte) bool {
    return pipe == 'S' || pipe == '|' || pipe == 'J' || pipe == 'L'
}

func fitsSouth(pipe byte) bool {
    return pipe == 'S' || pipe == '|' || pipe == 'F' || pipe == '7'
}

func locationAtCoordinates(pipesMap PipesMap, x int, y int) *Location {
	if x < 0 || x >= MapSize || y < 0 || y >= MapSize {
		return nil
	}

	return &Location{
		pipe: pipesMap[y][x],
		x:    x,
		y:    y,
	}
}
