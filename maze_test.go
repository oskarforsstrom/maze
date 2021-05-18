package maze

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g_1 := NewGraph(1, 1)
	g_5 := NewGraph(5, 5)
	g_100 := NewGraph(100, 100)

	var tests = []struct {
		vertices map[string]*vertex
		exp      int
	}{
		{g_1.vertices, 1},
		{g_5.vertices, 25},
		{g_100.vertices, 100},
	}
	for _, e := range tests {
		res := len(e.vertices)
		if res != e.exp {
			t.Errorf("len(Graph.vertices) = %v, expected: %v", res, e.exp)
		}
	}

	for _, vert := range g_5.vertices {
		y, x := coordToInt(vert.key)
		res := len(vert.neighbours)
		if x == 1 && y == 1 {
			if res != 2 {
				t.Errorf("len(%v.neighbours) = %v, expected: 2", vert.key, res)
			}
		} else if y == 1 && x == 5 {
			if res != 2 {
				t.Errorf("len(%v.neighbours) = %v, expected: 2", vert.key, res)
			}
		} else if y == 5 && x == 1 {
			if res != 2 {
				t.Errorf("len(%v.neighbours) = %v, expected: 2", vert.key, res)
			}
		} else if y == 5 && x == 5 {
			if res != 2 {
				t.Errorf("len(%v.neighbours) = %v, expected: 2", vert.key, res)
			}
		} else if x == 1 && res != 3 {
			t.Errorf("len(%v.neighbours) = %v, expected: 3", vert.key, res)
		} else if x == 5 && res != 3 {
			t.Errorf("len(%v.neighbours) = %v, expected: 3", vert.key, res)
		} else if y == 1 && res != 3 {
			t.Errorf("len(%v.neighbours) = %v, expected: 3", vert.key, res)
		} else if y == 5 && res != 3 {
			t.Errorf("len(%v.neighbours) = %v, expected: 3", vert.key, res)
		} else if res != 4 {
			t.Errorf("len(%v.neighbours) = %v, expected: 4", vert.key, res)
		}
	}

	h := g_5.height
	w := g_5.width
	if h != 5 {
		t.Errorf("g_5.heigth = %v, expected: 5", h)
	}
	if w != 5 {
		t.Errorf("g_5.width = %v, expected", w)
	}
}

func TestString(t *testing.T) {
	g_5 := NewGraph(5, 5)
	for i := 1; i <= 3; i++ {
		g_5.AddObstacle(i, 3)
	}

	g_1_2 := NewGraph(1, 2)

	var tests = []struct {
		g   Graph
		exp string
	}{
		{g_5, ".-------.-------.-------.-------.-------.\n| ( s )   (1,2) | (1,3) | (1,4)   ( f ) |\n:       +       +-------+       +       +\n| (2,1)   (2,2) | (2,3) | (2,4)   (2,5) |\n:       +       +-------+       +       +\n| (3,1)   (3,2) | (3,3) | (3,4)   (3,5) |\n:       +       +-------+       +       +\n| (4,1)   (4,2)   (4,3)   (4,4)   (4,5) |\n:       +       +       +       +       +\n| (5,1)   (5,2)   (5,3)   (5,4)   (5,5) |\n'-------'-------'-------'-------'-------'"},
		{g_1_2, ".-------.-------.\n| (1,1)   (1,2) |\n'-------'-------'"},
	}
	for _, e := range tests {
		res := e.g.String()
		if res != e.exp {
			t.Errorf("g.String() = %v, expected: %v", res, e.exp)
		}
	}
}

func TestAddObstacle(t *testing.T) {
	g_5 := NewGraph(5, 5)
	for i := 1; i <= 3; i++ {
		g_5.AddObstacle(i, 3)
	}
	for i := 1; i <= 3; i++ {
		if !g_5.vertices[coordinate(i, 3)].obstacle {
			t.Errorf("Error: AddObstacle not working properly, vertex.obstacle.")
		}
		if len(g_5.vertices[coordinate(i, 3)].neighbours) != 0 {
			t.Errorf("Error: AddObstacle not working properly, vertex.neighbours.")
		}
		switch i {
		case 1:
			if (len(g_5.vertices[coordinate(i, 2)].neighbours) != 2) || (len(g_5.vertices[coordinate(i, 4)].neighbours) != 2) {
				t.Errorf("Error: AddObstacle not working properly, vertex.neighbours.")
			}
		default:
			if (len(g_5.vertices[coordinate(i, 2)].neighbours) != 3) || (len(g_5.vertices[coordinate(i, 4)].neighbours) != 3) {
				t.Errorf("Error: AddObstacle not working properly, vertex.neighbours.")
			}
		}
		if len(g_5.vertices[coordinate(3, 4)].neighbours) != 3 {
			t.Errorf("Error: AddObstacle not working properly, vertex.neighbours.")
		}
	}
}

func TestRemoveObstacle(t *testing.T) {
	g_3 := NewGraph(3, 3)
	g_3.AddObstacle(2, 2)
	g_3.AddObstacle(1, 2)
	g_3.RemoveObstacle(2, 2)
	g_3.RemoveObstacle(3, 3) // RemoveObstacle on vertex that isn't an obstacle
	if g_3.vertices[coordinate(2, 2)].obstacle {
		t.Errorf("Error: RemoveObstacle not working properly, vertex.obstacle.")
	}
	if len(g_3.vertices[coordinate(2, 2)].neighbours) != 3 {
		t.Errorf("Error: RemoveObstacle not working properly, vertex.neighbours.")
	}
	if len(g_3.vertices[coordinate(1, 2)].neighbours) != 0 {
		t.Errorf("Error: RemoveObstacle not working properly, vertex.neighbours.")
	}
}

func TestAddStart(t *testing.T) {
	g := NewGraph(5, 5)
	g.AddStart(1, 1)
	for _, v := range g.vertices {
		if v.key == coordinate(1, 1) {
			if !v.startVertex {
				t.Errorf("Error: AddStart not working properly, vertex.startVertex.")
			}
		} else {
			if v.startVertex {
				t.Errorf("Error: AddStart not working properly, vertex.startVertex.")
			}
		}
	}
	g.AddStart(3, 4)
	for _, v := range g.vertices {
		if v.key == coordinate(3, 4) {
			if !v.startVertex {
				t.Errorf("Error: AddStart not working properly, vertex.startVertex.")
			}
		} else {
			if v.startVertex {
				t.Errorf("Error: AddStart not working properly, vertex.startVertex.")
			}
		}
	}
	if g.start != coordinate(3, 4) {
		t.Errorf("Error: AddStart not working properly, g.start")
	}
}

func TestAddFinish(t *testing.T) {
	g := NewGraph(5, 5)
	g.AddFinish(1, 1)
	for _, v := range g.vertices {
		if v.key == coordinate(1, 1) {
			if !v.finishVertex {
				t.Errorf("Error: AddFinish not working properly, vertex.finishVertex.")
			}
		} else {
			if v.finishVertex {
				t.Errorf("Error: AddFinish not working properly, vertex.finishVertex.")
			}
		}
	}
	g.AddFinish(3, 4)
	for _, v := range g.vertices {
		if v.key == coordinate(3, 4) {
			if !v.finishVertex {
				t.Errorf("Error: AddFinish not working properly, vertex.finishVertex.")
			}
		} else {
			if v.finishVertex {
				t.Errorf("Error: AddFinish not working properly, vertex.finishVertex.")
			}
		}
	}
	if g.finish != coordinate(3, 4) {
		t.Errorf("Error: AddFinish not working properly, g.finish")
	}
}

func TestStringFastestPath(t *testing.T) {
	g := NewGraph(5, 3)
	g.AddStart(1, 1)
	g.AddFinish(1, 3)
	g.AddObstacle(1, 2)
	g.AddObstacle(2, 2)
	g.AddObstacle(4, 2)
	if g.StringFastestPath() != "\n.-------.-------.-------.\n| ( s ) | (1,2) | ( f ) |\n:       +-------+       +\n| ( p ) | (2,2) | ( p ) |\n:       +-------+       +\n| ( p )   ( p )   ( p ) |\n:       +-------+       +\n| (4,1) | (4,2) | (4,3) |\n:       +-------+       +\n| (5,1)   (5,2)   (5,3) |\n'-------'-------'-------'\n\ndistance =6" {
		t.Errorf("Error: StringFastestPath not working properly.")
	}
}

func TestGetFastestPath(t *testing.T) {
	g := NewGraph(5, 3)
	g.AddStart(1, 1)
	g.AddFinish(1, 3)
	g.AddObstacle(1, 2)
	g.AddObstacle(2, 2)
	g.AddObstacle(4, 2)
	exp := []string{
		"(1,1)",
		"(2,1)",
		"(3,1)",
		"(3,2)",
		"(3,3)",
		"(2,3)",
		"(1,3)",
	}
	i, s := g.GetFastestPath()
	if i != 6 || !stringSliceEq(s, exp) {
		t.Errorf("g.GetFastestPath() = %v, %v; expected: %v, %v", i, s, 6, exp)
	}
}

func stringSliceEq(slice1 []string, slice2 []string) bool {
	if (slice1 == nil) != (slice2 == nil) {
		return false
	}
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
