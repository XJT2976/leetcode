package leetcode200

func NumIslands(grid [][]byte) int {
	num := 0
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				bfs(i, j, m, n, grid, visited)
				num++
			}
		}
	}

	return num
}

func bfs(x, y, m, n int, grid [][]byte, visited [][]bool) {
	queue := make([][2]int, 0)
	visited[x][y] = true
	queue = append(queue, [2]int{x, y})
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top[1]+1 < n && !visited[top[0]][top[1]+1] && grid[top[0]][top[1]+1] == '1' {
			visited[top[0]][top[1]+1] = true
			queue = append(queue, [2]int{top[0], top[1] + 1})
		}
		if top[1]-1 >= 0 && !visited[top[0]][top[1]-1] && grid[top[0]][top[1]-1] == '1' {
			visited[top[0]][top[1]-1] = true
			queue = append(queue, [2]int{top[0], top[1] - 1})
		}
		if top[0]+1 < m && !visited[top[0]+1][top[1]] && grid[top[0]+1][top[1]] == '1' {
			visited[top[0]+1][top[1]] = true
			queue = append(queue, [2]int{top[0] + 1, top[1]})
		}
		if top[0]-1 >= 0 && !visited[top[0]-1][top[1]] && grid[top[0]-1][top[1]] == '1' {
			visited[top[0]-1][top[1]] = true
			queue = append(queue, [2]int{top[0] - 1, top[1]})
		}
	}
}
