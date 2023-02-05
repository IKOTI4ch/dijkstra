package algo

import (
	"dijkstra/pkg/log"
	"dijkstra/pkg/vertices"
	"sort"
)

// ShortestAll calculates all the shortest paths from src to dest
func (g *Graph) ShortestAll(src, dest int) (BestPaths, error) {
	return g.evaluateAll(src, dest, 1, true)
}

// LimitedShortestAll calculates all the limited shortest paths from src to dest
func (g *Graph) LimitedShortestAll(src, dest int, limit int64) (BestPaths, error) {
	return g.evaluateAll(src, dest, limit, true)
}

// LongestAll calculates all the longest paths from src to dest
func (g *Graph) LongestAll(src, dest int) (BestPaths, error) {
	return g.evaluateAll(src, dest, 1, false)
}

// LimitedLongestAll calculates all the limited longest paths from src to dest
func (g *Graph) LimitedLongestAll(src, dest int, limit int64) (BestPaths, error) {
	return g.evaluateAll(src, dest, limit, false)
}

func (g *Graph) evaluateAll(src, dest int, limit int64, shortest bool) (BestPaths, error) {
	// Setup graph
	g.setup(shortest, src, -1)

	return g.postSetupEvaluateAll(src, dest, limit, shortest)
}

func (g *Graph) postSetupEvaluateAll(src, dest int, limit int64, shortest bool) (BestPaths, error) {
	var current *vertices.Vertex
	oldCurrent := -1

	for g.visiting.Len() > 0 {
		// Visit the current lowest distanced vertex
		current = g.visiting.PopOrdered()
		if oldCurrent == current.ID {
			continue
		}

		oldCurrent = current.ID
		// If the current Distance is already worse
		// than the best try another vertices.Vertex
		if shortest && current.Distance > g.Best {
			continue
		}

		for v, dist := range current.Arcs {
			// If the arc has better access, than the current best, update the vertices.Vertex being touched
			//if (shortest && current.Distance+dist < g.Vertices[v].Distance) ||
			if (shortest && current.Distance+dist < g.Vertices[v].Distance) ||
				(!shortest && current.Distance+dist > g.Vertices[v].Distance) ||
				(current.Distance+dist == g.Vertices[v].Distance && !g.Vertices[v].ContainsBest(current.ID)) ||
				(g.Best < limit && !current.ContainsBest(v)) { //TODO
				// if g.Vertices[v].bestVertex == current.ID && g.Vertices[v].ID != dest {
				if current.ContainsBest(v) && g.Vertices[v].ID != dest {
					// also only do this if we aren't checkout out the best Distance again
					return BestPaths{}, log.NewErrLoop(current.ID, v)
				}

				if current.Distance+dist == g.Vertices[v].Distance ||
					((current.ID == dest || v == dest) && (g.Vertices[dest].Distance <= limit)) {
					//At this point we know it's not in the list due to initial check
					g.Vertices[v].BestVertices = append(g.Vertices[v].BestVertices, current.ID)
				} else {
					g.Vertices[v].Distance = current.Distance + dist
					g.Vertices[v].BestVertices = []int{current.ID}
				}

				if v == dest {
					g.VisitedTarget = true
					g.Best = current.Distance + dist
					continue
					// If this is the destination update best, so we can stop looking at
					// useless Vertices
				}
				// Push this updated vertices.Vertex into the
				// list to be evaluated, pushes in sorted form
				g.visiting.PushOrdered(&g.Vertices[v])
			}
		}
	}

	if !g.VisitedTarget {
		return BestPaths{}, log.ErrNoPath
	}

	return g.bestPaths(src, dest, int(limit)), nil
}

func (g *Graph) bestPaths(src, dest, limit int) BestPaths {
	paths := g.visitPath(src, dest, dest)
	best := BestPaths{}

	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	for indexPaths, p := range paths {
		//if len(p) == limit {
		//	break
		//} //TODO

		for i, j := 0, len(paths[indexPaths])-1; i < j; i, j = i+1, j-1 {
			paths[indexPaths][i], paths[indexPaths][j] = paths[indexPaths][j], paths[indexPaths][i]
		}

		best = append(best, BestPath{int64(len(p) - 1), paths[indexPaths]})
	}

	return best
}

func (g *Graph) visitPath(src, dest, currentNode int) [][]int {
	if currentNode == src {
		return [][]int{
			{currentNode},
		}
	}

	paths := make([][]int, 0)
	for _, vertex := range g.Vertices[currentNode].BestVertices {
		sps := g.visitPath(src, dest, vertex)

		for i := range sps {
			paths = append(paths, append([]int{currentNode}, sps[i]...))
		}
	}

	return paths
}
