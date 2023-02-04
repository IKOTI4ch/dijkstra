package algo

import (
	"dijkstra/pkg/vertices"
	"errors"
)

type Graph struct {
	Best            int64
	VisitedTarget   bool
	Vertices        vertices.Vertices // slice of all vertices available
	visiting        queue
	mapping         map[string]int
	usingMap        bool
	highestMapIndex int
}

// NewGraph creates a new empty algo
func NewGraph() *Graph {
	newGraph := &Graph{}
	newGraph.mapping = map[string]int{}
	return newGraph
}

// AddNewVertex adds a new vertex at the next available index
func (g *Graph) AddNewVertex() *vertices.Vertex {
	for i, v := range g.Vertices {
		if i != v.ID {
			g.Vertices[i] = vertices.Vertex{ID: i}
			return &g.Vertices[i]
		}
	}

	return g.AddVertex(len(g.Vertices))
}

// AddVertex adds a single vertex
func (g *Graph) AddVertex(ID int) *vertices.Vertex {
	g.AddVertices(vertices.Vertex{ID: ID})
	return &g.Vertices[ID]
}

// AddVertices adds the listed vertices to the algo, overwrites any existing
// Vertex with the same ID.
func (g *Graph) AddVertices(vs ...vertices.Vertex) {
	for _, v := range vs {
		v.BestVertices = []int{-1}
		if v.ID >= len(g.Vertices) {
			newN := make(vertices.Vertices, v.ID+1-len(g.Vertices))
			g.Vertices = append(g.Vertices, newN...)
		}
		g.Vertices[v.ID] = v
	}
}

// GetVertex gets the reference of the specified vertex.
// An error is thrown if there is no vertex with that index/ID.
func (g *Graph) GetVertex(ID int) (*vertices.Vertex, error) {
	if ID >= len(g.Vertices) {
		return nil, errors.New("vertex not found")
	}
	return &g.Vertices[ID], nil
}

// SetDefaults sets the distance and best vertex to that specified
func (g *Graph) setDefaults(Distance int64, BestVertex int) {
	for i := range g.Vertices {
		g.Vertices[i].BestVertices = []int{BestVertex}
		g.Vertices[i].Distance = Distance
	}
}
