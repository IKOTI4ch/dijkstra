package algo

import (
	"dijkstra/pkg/log"
	"dijkstra/pkg/vertices"
	"math"
)

// Shortest calculates the shortest path from src to dest
func (g *Graph) Shortest(src, dest int) (BestPath, error) {
	return g.evaluate(src, dest, true)
}

// Longest calculates the longest path from src to dest
func (g *Graph) Longest(src, dest int) (BestPath, error) {
	return g.evaluate(src, dest, false)
}

func (g *Graph) setup(shortest bool, src int, list int) {
	//-1 auto list
	//Get a new list regardless
	if list >= 0 {
		g.forceList(list)
	} else if shortest {
		g.forceList(-1)
	} else {
		g.forceList(-2)
	}

	//Reset state
	g.VisitedTarget = false
	//Reset the best current value (the worst, so it gets overwritten)
	// and set the defaults *almost* as bad
	// set all best vertices to -1 (unused)
	if shortest {
		g.setDefaults(int64(math.MaxInt64)-2, -1)
		g.Best = int64(math.MaxInt64)
	} else {
		g.setDefaults(int64(math.MinInt64)+2, -1)
		g.Best = int64(math.MinInt64)
	}

	//Set the distance of initial vertex 0
	g.Vertices[src].Distance = 0
	//Add the source vertex to the list
	g.visiting.PushOrdered(&g.Vertices[src])
}

func (g *Graph) forceList(i int) {
	//-2 long auto
	//-1 short auto
	//0 short pq
	//1 long pq
	//2 short ll
	//3 long ll
	switch i {
	case -2:
		if len(g.Vertices) < 800 {
			g.forceList(2)
		} else {
			g.forceList(0)
		}
		break
	case -1:
		if len(g.Vertices) < 800 {
			g.forceList(3)
		} else {
			g.forceList(1)
		}
		break
	case 0:
		g.visiting = priorityQueueNewShort()
		break
	case 1:
		g.visiting = priorityQueueNewLong()
		break
	case 2:
		g.visiting = linkedListNewShort()
		break
	case 3:
		g.visiting = linkedListNewLong()
		break
	default:
		panic(i)
	}
}

func (g *Graph) bestPath(src, dest int) BestPath {
	var path []int

	for c := g.Vertices[dest]; c.ID != src; c = g.Vertices[c.BestVertices[0]] {
		path = append(path, c.ID)
	}
	path = append(path, src)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return BestPath{g.Vertices[dest].Distance, path}
}

func (g *Graph) evaluate(src, dest int, shortest bool) (BestPath, error) {
	//Setup algo
	g.setup(shortest, src, -1)

	return g.postSetupEvaluate(src, dest, shortest)
}

func (g *Graph) postSetupEvaluate(src, dest int, shortest bool) (BestPath, error) {
	var current *vertices.Vertex
	oldCurrent := -1

	for g.visiting.Len() > 0 {
		//Visit the current lowest distanced Vertex
		current = g.visiting.PopOrdered()

		if oldCurrent == current.ID {
			continue
		}

		oldCurrent = current.ID
		//If the current distance is already worse than the best try another Vertex
		if shortest && current.Distance >= g.Best {
			continue
		}

		for v, dist := range current.Arcs {
			//If the arc has better access, than the current best, update the Vertex being touched
			if (shortest && current.Distance+dist < g.Vertices[v].Distance) ||
				(!shortest && current.Distance+dist > g.Vertices[v].Distance) {
				if current.BestVertices[0] == v && g.Vertices[v].ID != dest {
					//also only do this if we aren't checkout out the best distance again
					//This seems familiar 8^)
					return BestPath{}, log.NewErrLoop(current.ID, v)
				}
				g.Vertices[v].Distance = current.Distance + dist
				g.Vertices[v].BestVertices[0] = current.ID
				if v == dest {
					//If this is the destination update best, so we can stop looking at
					// useless Vertices
					g.Best = current.Distance + dist
					g.VisitedTarget = true
					continue // Do not push if dest
				}
				//Push this updated Vertex into the list to be evaluated, pushes in
				// sorted form
				g.visiting.PushOrdered(&g.Vertices[v])
			}
		}
	}

	return g.finally(src, dest)
}

func (g *Graph) finally(src, dest int) (BestPath, error) {
	if !g.VisitedTarget {
		return BestPath{}, log.ErrNoPath
	}

	return g.bestPath(src, dest), nil
}
