package sol

type Pair struct {
	row, col int
}

func numIslands(grid [][]byte) int {
	ROW := len(grid)
	COL := len(grid[0])
	visit := make([][]bool, ROW)
	for row := range visit {
		visit[row] = make([]bool, COL)
		for col := range visit[row] {
			visit[row][col] = false
		}
	}
	island := 0
	directions := []Pair{{row: -1, col: 0}, {row: 1, col: 0}, {row: 0, col: -1}, {row: 0, col: 1}}
	var bfs = func(row int, col int) {
		queue := []Pair{{row: row, col: col}}
		visit[row][col] = true
		for len(queue) != 0 {
			top := queue[0]
			queue = queue[1:]
			for _, direction := range directions {
				shifted_row := top.row + direction.row
				shifted_col := top.col + direction.col
				if shifted_row < 0 || shifted_row >= ROW ||
					shifted_col < 0 || shifted_col >= COL ||
					visit[shifted_row][shifted_col] || grid[shifted_row][shifted_col] == '0' {
					continue
				}
				visit[shifted_row][shifted_col] = true
				queue = append(queue, Pair{row: shifted_row, col: shifted_col})
			}
		}
	}
	for row := 0; row < ROW; row++ {
		for col := 0; col < COL; col++ {
			if grid[row][col] == '1' && !visit[row][col] {
				bfs(row, col)
				island++
			}
		}
	}
	return island
}
