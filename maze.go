// Oskar Forsström, grudat 21 ovn 7
//
// Package graphTraversal provides a unweighted graph structure with start- and finishvertices
// and a way of finding the shortest path from start to finish.
//
// When a maze has multiple solutions one might wanna find the shortest possible path from start
// to finish. The BFS method of finding the shortest possible path uses a queue to visit vertice
// with increasing distance from the start-vertex and thereby finds the shortest possible path.
//
package maze

import (
	"fmt"
	"strconv"
	"strings"
)

// Graph represents a maze-graph.
//
// A graph where every vertex is represented by a 2D coordinate
// (a unique combination of two integers) and where every vertex only
// can have edges between itself and adjencent (non-diagonal) vertices,
// i.e., if we're given a vertices (x,y) we know that it can only have edges to the vertices:
// (x+1,y), (x-1,y), (x,y+1), (x,y-1)
type Graph struct {
	// Width and Height are integers representing the graph's size.
	//
	// For example if Width = n, Heigth = m, the graph will look like:
	//
	// .-------.-------.-------.-----.-------.
	// | (1,1) | (1,2) | (1,3) | ... | (1,n) |
	// :-------+-------+-------+-----+-------:
	// | (2,1) | (2,2) | (2,3) |     |       |
	// :-------+-------+-------+-----+-------:
	// | (3,1) | (3,2) | (3,3) |     |       |
	// :-------+-------+-------+-----+-------:
	// | ...   |       |       | ... |       |
	// :-------+-------+-------+-----+-------:
	// | (m,1) |       |       |     | (m,n) |
	// '-------'-------'-------'-----'-------'
	height int
	width  int

	// Start and finish contains the unique keys to the start- and finishVertex.
	start  string
	finish string

	// vertices is a map containing pointers to all vertices in the graph.
	vertices map[string]*vertex
}

// vertex is the object for every vertex in the graph.
type vertex struct {
	// key is the unique representation to every vertex in the graph.
	// Every key has the format "(y,x)", where y and x represents its
	// coordinates in the graph.
	key string

	// start- and finishVertex is a boolean signifying if the vertex is
	// a start-/finishvertex or not.
	startVertex  bool
	finishVertex bool

	// obstacle is a boolean signifying if the vertex is a
	// obstacle or not.
	obstacle bool

	// visited is a boolean signifying if the vertex has been visited or not.
	visited bool

	// neighbours is a map of pointers to the vertices who have edges connected to
	// the given vertex. A vertex can only have edges to adjencent vertices (non-diagonal),
	// i.e., if we're given a vertex (x,y) we know that it can only have edges to the vertices:
	// (x+1,y), (x-1,y), (x,y+1), (x,y-1)
	// [Given that all Coordinates are within the heigth and width specifications]
	neighbours map[string]*vertex
}

// NewGraph creates a graph of size, width x Heigth, where every vertex has a edge connected
// to every adjencent vertex (non-diagonal).
func NewGraph(height int, width int) Graph {
	if width <= 0 || height <= 0 {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("Error: width or height <= 0")
	}
	var graph Graph
	graph.height = height
	graph.width = width
	graph.vertices = make(map[string]*vertex)

	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			coord := coordinate(i, j)
			var vert vertex
			vert.key = coord
			graph.vertices[coord] = &vert
		}
	}
	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			coord := coordinate(i, j)
			neighbours := make(map[string]*vertex)
			switch i {
			case 1:
				neighbours[coordinate(i+1, j)] = graph.vertices[coordinate(i+1, j)]
			case height:
				neighbours[coordinate(i-1, j)] = graph.vertices[coordinate(i-1, j)]
			default:
				neighbours[coordinate(i+1, j)] = graph.vertices[coordinate(i+1, j)]
				neighbours[coordinate(i-1, j)] = graph.vertices[coordinate(i-1, j)]
			}
			switch j {
			case 1:
				neighbours[coordinate(i, j+1)] = graph.vertices[coordinate(i, j+1)]
			case width:
				neighbours[coordinate(i, j-1)] = graph.vertices[coordinate(i, j-1)]
			default:
				neighbours[coordinate(i, j+1)] = graph.vertices[coordinate(i, j+1)]
				neighbours[coordinate(i, j-1)] = graph.vertices[coordinate(i, j-1)]
			}
			graph.vertices[coord].neighbours = neighbours
		}
	}
	return graph
}

// coordinate takes two integers and returns a string representation
// of the two integers as a coordinate, on the format:
// "(y,x)"
func coordinate(y int, x int) string {
	return "(" + strconv.Itoa(y) + "," + strconv.Itoa(x) + ")"
}

// The method String returns a string ASCII representation of the graph
// with visual representation for vertices, edges, startVertex
// and finishVertex.
func (g *Graph) String() string {
	stringGrid := make([][]string, 2*g.height+1)
	for i := 0; i <= 2*g.height; i++ {
		stringGrid[i] = make([]string, g.width+1)
	}

	stringGrid[0][0] = "."
	for i := 2; i <= 2*g.height-2; i += 2 {
		stringGrid[i][0] = ":"
	}
	for i := 1; i <= 2*g.height-1; i += 2 {
		stringGrid[i][0] = "|"
	}
	stringGrid[2*g.height][0] = "'"
	for j := 1; j <= g.width; j++ {
		stringGrid[0][j] = "-------."
		idx := 1
		for i := 2; i <= 2*g.height-2; i += 2 {
			vertex := g.vertices[coordinate(idx, j)]
			_, found := vertex.neighbours[coordinate(idx+1, j)]
			if found {
				stringGrid[i][j] = "       +"
			} else {
				stringGrid[i][j] = "-------+"
			}
			idx++
		}
		idx = 1
		for i := 1; i <= 2*g.height-1; i += 2 {
			coord := coordinate(idx, j)
			if g.vertices[coordinate(idx, j)].startVertex {
				coord = "( s )"
			}
			if g.vertices[coordinate(idx, j)].finishVertex {
				coord = "( f )"
			}
			vertex := g.vertices[coordinate(idx, j)]
			_, found := vertex.neighbours[coordinate(idx, j+1)]
			if found {
				stringGrid[i][j] = " " + coord + "  "
			} else {
				stringGrid[i][j] = " " + coord + " |"
			}
			idx++
		}
		stringGrid[2*g.height][j] = "-------'"
	}
	var result string
	for i := 0; i <= 2*g.height; i++ {
		result += strings.Join(stringGrid[i], "") + "\n"
	}
	return result
}

// The method AddObstacle adds an obstacle at the specified vertex,
// i.e. removes edges between the vertex and its adjencent vertices
// and changes vertex.obstacle to true.
func (g *Graph) AddObstacle(y int, x int) {
	if g.vertices[coordinate(y, x)].startVertex || g.vertices[coordinate(y, x)].finishVertex {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("Error: the specified vertex is a start- or finsihVertex")
	}
	switch y {
	case 1:
		g.removeEdge(coordinate(y, x), coordinate(y+1, x))
	case g.height:
		g.removeEdge(coordinate(y, x), coordinate(y-1, x))
	default:
		g.removeEdge(coordinate(y, x), coordinate(y+1, x))
		g.removeEdge(coordinate(y, x), coordinate(y-1, x))
	}
	switch x {
	case 1:
		g.removeEdge(coordinate(y, x), coordinate(y, x+1))
	case g.width:
		g.removeEdge(coordinate(y, x), coordinate(y, x-1))
	default:
		g.removeEdge(coordinate(y, x), coordinate(y, x+1))
		g.removeEdge(coordinate(y, x), coordinate(y, x-1))
	}
	g.vertices[coordinate(y, x)].obstacle = true
}
func (g *Graph) removeEdge(coord1 string, coord2 string) {
	delete(g.vertices[coord1].neighbours, coord2)
	delete(g.vertices[coord2].neighbours, coord1)
}

// The method RemoveObstacle removes an obstacle at the specified vertex,
// i.e. adds edges between the vertex and its adjencent vertices, if the
// adjencent vertex isn't an obstacle. And changes vertex.obstacle to false.
func (g *Graph) RemoveObstacle(y int, x int) {
	switch y {
	case 1:
		if !g.vertices[coordinate(y+1, x)].obstacle {
			g.addEdge(coordinate(y, x), coordinate(y+1, x))
		}
	case g.height:
		if !g.vertices[coordinate(y-1, x)].obstacle {
			g.addEdge(coordinate(y, x), coordinate(y-1, x))
		}
	default:
		if !g.vertices[coordinate(y+1, x)].obstacle {
			g.addEdge(coordinate(y, x), coordinate(y+1, x))
		}
		if !g.vertices[coordinate(y-1, x)].obstacle {
			g.addEdge(coordinate(y, x), coordinate(y-1, x))
		}
	}
	switch x {
	case 1:
		if !g.vertices[coordinate(y, x+1)].obstacle {
			g.addEdge(coordinate(y, x), coordinate(y, x+1))
		}
	case g.width:
		if !g.vertices[coordinate(y, x-1)].obstacle {
			g.addEdge(coordinate(y, x), coordinate(y, x-1))
		}
	default:
		if !g.vertices[coordinate(y, x+1)].obstacle {
			g.addEdge(coordinate(y, x), coordinate(y, x+1))
		}
		if !g.vertices[coordinate(y, x-1)].obstacle {
			g.addEdge(coordinate(y, x), coordinate(y, x-1))
		}
	}
	g.vertices[coordinate(y, x)].obstacle = false
}
func (g *Graph) addEdge(coord1 string, coord2 string) {
	g.vertices[coord1].neighbours[coord2] = g.vertices[coord2]
	g.vertices[coord2].neighbours[coord1] = g.vertices[coord1]
}

// The method AddStart marks the specified vertex as the "startVertex".
//
// Does this by changing its field startVertex to true, if the specified vertex
// isn't an obstacle. If there already exists a startVertex, the method
// also turns the old startvertex into a "normal vertex".
//
// The method updatedsthe field Start of the Graph to match the new
// startVertex.
func (g *Graph) AddStart(y int, x int) {
	coord := coordinate(y, x)
	if g.vertices[coord].obstacle {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("Error: the specified vertex is an obstacle")
	}
	if g.vertices[coord].finishVertex {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("Error: the specified vertex is a finishVertex")
	}
	if g.start != "" {
		g.vertices[g.start].startVertex = false
	}
	g.vertices[coord].startVertex = true
	g.start = coord
}

// The method AddFinish marks the specified vertex as the "finishVertex".
//
// Does this by changing its field finishVertex to true, if the specified vertex
// isn't an obstacle. If there already exists a finishVertex, the method
// also turns the old finishVertex into a "normal vertex".
//
// The method updateds the field finish of the Graph to match the new
// finishVertex.
func (g *Graph) AddFinish(y int, x int) {
	coord := coordinate(y, x)
	if g.vertices[coord].obstacle {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("Error: the specified vertex is an obstacle")
	}
	if g.vertices[coord].startVertex {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("Error: the specified vertex is a startVertex")
	}
	if g.finish != "" {
		g.vertices[g.finish].finishVertex = false
	}
	g.vertices[coord].finishVertex = true
	g.finish = coord
}

func (g *Graph) unmarkVisited() {
	for i := 1; i <= g.height; i++ {
		for j := 1; j <= g.width; j++ {
			coord := coordinate(i, j)
			g.vertices[coord].visited = false
		}
	}
}

// The method StringFastestPath returns a ASCII representation of the shortest
// path between the start- and finishvertex and the distance of the path.
//
// The representation displays
// the start-vertex as: ( s )
// the finish-vertex as: ( f )
// the path as: ( p )
func (g *Graph) StringFastestPath() string {
	distance, path := g.GetFastestPath()

	stringGrid := make([][]string, 2*g.height+1)
	for i := 0; i <= 2*g.height; i++ {
		stringGrid[i] = make([]string, g.width+1)
	}

	stringGrid[0][0] = "\n."
	for i := 2; i <= 2*g.height-2; i += 2 {
		stringGrid[i][0] = ":"
	}
	for i := 1; i <= 2*g.height-1; i += 2 {
		stringGrid[i][0] = "|"
	}
	stringGrid[2*g.height][0] = "'"
	for j := 1; j <= g.width; j++ {
		stringGrid[0][j] = "-------."
		idx := 1
		for i := 2; i <= 2*g.height-2; i += 2 {
			vertex := g.vertices[coordinate(idx, j)]
			_, found := vertex.neighbours[coordinate(idx+1, j)]
			if found {
				stringGrid[i][j] = "       +"
			} else {
				stringGrid[i][j] = "-------+"
			}
			idx++
		}
		idx = 1
		for i := 1; i <= 2*g.height-1; i += 2 {
			coord := coordinate(idx, j)
			if g.vertices[coordinate(idx, j)].startVertex {
				coord = "( s )"
			}
			if g.vertices[coordinate(idx, j)].finishVertex {
				coord = "( f )"
			}
			vertex := g.vertices[coordinate(idx, j)]
			_, found := vertex.neighbours[coordinate(idx, j+1)]
			if found {
				stringGrid[i][j] = " " + coord + "  "
			} else {
				stringGrid[i][j] = " " + coord + " |"
			}
			idx++
		}
		stringGrid[2*g.height][j] = "-------'"
	}
	for _, coord := range path {
		if g.vertices[coord].startVertex || g.vertices[coord].finishVertex {
			// do nothing
		} else {
			y, x := coordToInt(coord)
			y_grid := 2*y - 1
			stringGrid[y_grid][x] = makeP(stringGrid[y_grid][x])
		}
	}
	var result string
	for i := 0; i <= 2*g.height; i++ {
		result += strings.Join(stringGrid[i], "") + "\n"
	}
	return result + "\n" + "distance =" + strconv.Itoa(distance)
}

func makeP(text string) string {
	slice := make([]string, 8)
	for idx, x := range text {
		slice[idx] = string(x)
	}
	slice[2] = " "
	slice[3] = "p"
	slice[4] = " "
	return strings.Join(slice, "")
}

func coordToInt(coord string) (int, int) {
	slice := make([]string, 5)
	for idx, x := range coord {
		slice[idx] = string(x)
	}
	result1, _ := strconv.Atoi(slice[1])
	result2, _ := strconv.Atoi(slice[3])
	return result1, result2
}

// The method GetFastestPath returns the shortest distance between the start- and finishvertex
// and a slice of strings representing the shortest path.
//
// The slice is on the format:
// [(1,1), (2,1), ..., (5,4), (5,5)]
func (g *Graph) GetFastestPath() (int, []string) {
	distance, predecessor := g.fastestPathBFS()
	dist := distance[g.finish]
	stringSlice := make([]string, dist+1)
	stringSlice[dist] = g.finish
	stringSlice[0] = g.start
	vertex := g.finish
	for j := dist - 1; j >= 1; j-- {
		stringSlice[j] = predecessor[vertex]
		vertex = predecessor[vertex]
	}
	return dist, stringSlice
}

func (g *Graph) fastestPathBFS() (map[string]int, map[string]string) {
	g.unmarkVisited()
	var queue []*vertex
	var a *vertex
	distance := make(map[string]int)
	predecessor := make(map[string]string)

	g.vertices[g.start].visited = true
	queue = append(queue, g.vertices[g.start])
	distance[g.start] = 0

	for len(queue) > 0 {
		a = queue[0]
		queue = queue[1:]
		for _, x := range a.neighbours {
			if !x.visited {
				x.visited = true
				distance[x.key] = distance[a.key] + 1
				predecessor[x.key] = a.key
				queue = append(queue, x)

				if x.finishVertex {
					return distance, predecessor
				}
			}
		}
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("Error: there's no path between the start- and finishVertex")
}
