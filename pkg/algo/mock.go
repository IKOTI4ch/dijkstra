package algo

func GetMockData() (nodes map[string]int, edges [][2]string) {
	nodes = make(map[string]int)
	edges = make([][2]string, 0)

	nodes["A"] = 0
	nodes["B"] = 1
	nodes["C"] = 2
	nodes["D"] = 3
	nodes["E"] = 4

	edges = append(
		edges,
		[2]string{"A", "B"},
		[2]string{"A", "C"},
		[2]string{"B", "C"},
		[2]string{"B", "D"},
		[2]string{"B", "E"},
		[2]string{"C", "D"},
		[2]string{"E", "D"},
		[2]string{"D", "A"}, //For dev
	)

	return nodes, edges
}
