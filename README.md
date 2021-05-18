# Maze builder and BFS shortest-path finder

### Go maze data structure and ASCII visualizer with function for finding the shortest path from given start- and endpoints

When a maze has multiple solutions one might wanna find the shortest possible path from start to finish. The BFS method of finding the shortest possible path uses a queue to visit vertice with increasing distance from the start-vertex and thereby finds the shortest possible path.

---

# Documentation
## Overview
The package builds a maze based on a object-oriented graph with vertices and edges with adjencency lists. Every vertex has a unique two integer (coordinate) representation and can only have edges to adjencent vertices (non-diagonal).
## Examples

<details>
  <summary>Examples</summary>
  Find the shortest path between start and finish in a maze.
  
    package main
    
    import (
        "fmt"
        "gits-15.sys.kth.se/grudat21/oghvfo-ovn7"
    )

    func main() {
	    g := NewGraph(5, 5)
	    g.AddStart(1, 1)
	    g.AddFinish(1, 5)
	    for i := 1; i <= 3; i++ {
	    	g.AddObstacle(i, 3)
	    }
	    fmt.Printf(g.StringFastestPath())
    }

    ---
    Output:

    .-------.-------.-------.-------.-------.
    | ( s )   (1,2) | (1,3) | ( p )   ( f ) |
    :       +       +-------+       +       +
    | ( p )   (2,2) | (2,3) | ( p )   (2,5) |
    :       +       +-------+       +       +
    | ( p )   (3,2) | (3,3) | ( p )   (3,5) |
    :       +       +-------+       +       +
    | ( p )   ( p )   ( p )   ( p )   (4,5) |
    :       +       +       +       +       +
    | (5,1)   (5,2)   (5,3)   (5,4)   (5,5) |
    '-------'-------'-------'-------'-------'
    distance = 10
</details>



## Types

### type Graph
    type Graph struct {
        // containts unexported fields
    }

Graph represents a maze-graph where every vertex is represented by a 2D coordinate (a unique combination of two integers) and where every vertex only can have edges between itself and adjencent (non-diagonal) vertices, i.e., if we're given a Vertex (x,y) we know that it can only have edges to the vertices: (x+1,y), (x-1,y), (x,y+1), (x,y-1).

### func NewGraph
    func NewGraph(height int, width int) Graph
NewGraph creates a graph of size, width x Heigth, where every vertex has a edge connected to every adjencent Vertex (non-diagonal).

### func (*Graph) String
    func (g *Graph) String() string
Reurns a ASCII representation of the graph with visual representation for vertices, edges, startVertex and finishVertex.

### func (*Graph) AddObstacle
    func (g *Graph) AddObstacle(y int, x int)
Adds an obstacle at the specified vertex, i.e. removes edges between the vertex and its adjencent vertices.

### func (*Graph) RemoveObstacle
    func (g *Graph) RemoveObstacle(y int, x int)
Removes an obstacle at the specified vertex, i.e. adds edges between the vertex and its adjencent vertices, if the adjencent vertex isn't an obstacle.

### func (*Graph) AddStart
    func (g *Graph) AddStart(y int, x int)
Marks the specified vertex as the "startVertex".

### func (*Graph) AddFinish
    func (g *Graph) AddFinish(y int, x int)
Marks the specified vertex as the "finishVertex".

### func (*Graph) StringFastestPath
    func (g *Graph) StringFastestPath() string
Return a ASCII representation of the shortest path between the start- and finishvertex and the distance of the path.
### func (*Graph) GetFastestPath
    func (g *Graph) GetFastestPath() (int, []string)
Returns the shortest distance between the start- and finishvertex and a slice of strings representing the shortest path.
