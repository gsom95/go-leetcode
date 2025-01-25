package findeventualsafestates

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	visited := make([]int, n) // 0 = unvisited, 1 = visiting, 2 = safe
	var result []int

	for i := range n {
		if isSafe(graph, i, visited) {
			result = append(result, i)
		}
	}
	return result
}

func isSafe(graph [][]int, node int, visited []int) bool {
	if visited[node] > 0 {
		return visited[node] == 2
	}

	visited[node] = 1 // Mark as visiting
	for _, neighbor := range graph[node] {
		if !isSafe(graph, neighbor, visited) {
			return false
		}
	}
	visited[node] = 2 // Mark as safe
	return true
}
