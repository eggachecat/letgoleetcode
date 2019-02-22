package main

import "lc/solutions/search"

func main() {
	s := &search.SolutionsFor37At0{}
	// board := [][]byte{
	// 	[]byte("53..7...."),
	// 	[]byte("6..195..."),
	// 	[]byte(".98....6."),
	// 	[]byte("8...6...3"),
	// 	[]byte("4..8.3..1"),
	// 	[]byte("7...2...6"),
	// 	[]byte(".6....28."),
	// 	[]byte("...419..5"),
	// 	[]byte("....8..79"),
	// }
	board := [][]byte{
		[]byte("534678912"),
		[]byte("672195348"),
		[]byte("198342567"),
		[]byte("859761423"),
		[]byte("426853791"),
		[]byte("713924856"),
		[]byte("961537284"),
		[]byte("...419..5"),
		[]byte("....86179"),
	}
	s.SolveSudoku(board)
}
