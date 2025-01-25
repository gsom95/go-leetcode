package count_servers_that_communicate

func countServers(grid [][]int) int {
	var (
		connectedServers = 0
		rows             = len(grid)
		cols             = len(grid[0])
		serversPerRow    = make([]int, rows)
		serversPerCol    = make([]int, cols)
	)

	for row, rowElems := range grid {
		for col, elem := range rowElems {
			if elem == 0 {
				continue
			}

			serversPerRow[row]++
			serversPerCol[col]++
		}
	}

	for row, rowElems := range grid {
		for col, elem := range rowElems {
			if elem == 1 && (serversPerRow[row] > 1 || serversPerCol[col] > 1) {
				connectedServers++
			}
		}
	}

	return connectedServers
}
