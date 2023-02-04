package algo

func (g *Graph) CreateData(nodes map[string]int, edges [][2]string) {
	var err error

	i := 0 // index of node
	for id := range nodes {
		nodes[id] = i
		g.AddVertex(nodes[id])
		i++
	}

	for _, vertices := range edges {
		err = g.AddArc(nodes[vertices[0]], nodes[vertices[1]], 1)
		if err != nil {
			panic(err)
		}

		err = g.AddArc(nodes[vertices[1]], nodes[vertices[0]], 1)
		if err != nil {
			panic(err)
		}
	}
}
