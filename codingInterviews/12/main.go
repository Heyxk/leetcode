package main

func exist(board [][]byte, word string) bool {
	for i, m := 0, len(board); i < m; i++ {
		for j, n := 0, len(board[0]); j < n; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false

}

func dfs(board [][]byte, word string, i int, j int, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	board[i][j] = ' '
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) || dfs(board, word, i, j+1, k+1) || dfs(board, word, i, j-1, k+1)
	board[i][j] = word[k]
	return res
}
