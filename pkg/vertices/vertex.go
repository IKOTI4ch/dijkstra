package vertices

type Vertex struct {
	ID           int
	Arcs         map[int]int64 // A set of all weights to the vertices in the map
	BestVertices []int
	Distance     int64 // Best distance to the Vertex
}

// NewVertex creates a new vertex
func NewVertex(ID int) *Vertex {
	return &Vertex{ID: ID, BestVertices: []int{-1}, Arcs: map[int]int64{}}
}

func (v *Vertex) ContainsBest(id int) bool {
	for _, bestVertex := range v.BestVertices {
		if bestVertex == id {
			return true
		}
	}
	return false
}

// AddArc adds an arc to the vertex, it's up to the user to make sure this is used
// correctly, firstly ensuring to use before adding to algo, or to use referenced
// of the Vertex instead of a copy. Secondly, to ensure the destination is a valid
// Vertex in the algo. Note that AddArc will overwrite any existing distance set
// if there is already an arc set to Destination.
func (v *Vertex) AddArc(Destination int, Distance int64) {
	if v.Arcs == nil {
		v.Arcs = map[int]int64{}
	}
	v.Arcs[Destination] = Distance
}

// RemoveArc completely removes the arc to Destination (if it exists), this method will
// not error if Destination doesn't exist, only nop
func (v *Vertex) RemoveArc(Destination int) {
	delete(v.Arcs, Destination)
}

// GetArc gets the specified arc to Destination, bool is false if no arc found
func (v *Vertex) GetArc(Destination int) (distance int64, ok bool) {
	if v.Arcs == nil {
		return 0, false
	}
	distance, ok = v.Arcs[Destination] //TODO doesn't work on one line?

	return
}
