# golang_number_of_islands

Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return *the number of islands*.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

## Examples

**Example 1:**

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

```

**Example 2:**

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

```

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`.

## 解析

給定一個 2 D 二元矩陣 grid ， 每個元素 grid[r][c] 只有 1 或是 0 兩種值

1 代表是陸地，0 代表是 水

如果水平或是垂直相鄰位置都是 1 代表 cell屬於同一個 island

題目實作一個演算法從 grid 找出總共有幾個 island

直覺去思考，可以發現要找出有幾個 island 代表要去找有幾個相連的區塊

而要找出相連的區塊會透過 BFS 從該區塊相鄰為 1 的 cell 開始找尋直到遇到 0 或是 邊界

如下圖

![](https://i.imgur.com/XYBldEo.png)

會發現需要紀錄每個拜訪過的 grid，可以使用 hashSet

然後逐步每個 cell 檢查

## 程式碼

```go
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

```
## 困難點

1. 需要理解找出連結在一起 cell 要使用 BFS
2. 需要理解 BFS 實作

## Solve Point

- [x]  Understand what problem to solve
- [x]  Analysis Complexity