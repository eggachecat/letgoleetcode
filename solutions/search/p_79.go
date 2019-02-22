package search

import "log"

type SolutionsFor79 interface {
	Exist(board [][]byte, word string) bool
}

type SolutionsFor79At0 struct{}

func getChoices(board [][]byte, targetChar byte, duplicate [][]bool, point []int) [][]int {
	maxRow, maxCol := len(board), len(board[0])
	row, col := point[0], point[1]
	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	choices := [][]int{}
	for i := 0; i < 4; i++ {
		dir := dirs[i]
		rowNew, colNew := row+dir[0], col+dir[1]
		if rowNew >= 0 && rowNew < maxRow && colNew >= 0 && colNew < maxCol {
			// 方向合法
			log.Println(rowNew, colNew, maxRow, maxCol, "legal")
			if !duplicate[rowNew][colNew] {
				// 没被走过
				log.Println(rowNew, colNew, maxRow, maxCol, "available")
				log.Println(board[rowNew][colNew], targetChar)
				if board[rowNew][colNew] == targetChar {
					// 且等于目标
					choices = append(choices, []int{rowNew, colNew})
				}
			}
		}
	}
	return choices
}

// Exist blabla
// word search
// Choice: 方向
// Constraints: 不能重复 + 选的要和目标的一样
// Goal: 走出一个长度一样的字符串
func (s *SolutionsFor79At0) Exist(board [][]byte, word string) bool {

	var permuate func(board [][]byte, duplicate [][]bool, point []int, nChar int, targetStr string) bool
	permuate = func(board [][]byte, duplicate [][]bool, point []int, nChar int, targetStr string) bool {
		if nChar == len(targetStr)-1 {
			return true
		}
		targetChar := byte(targetStr[nChar+1])
		choices := getChoices(board, targetChar, duplicate, point)
		log.Println("targetChar", targetChar, "choices", choices)

		for _, choice := range choices {
			duplicate[choice[0]][choice[1]] = true
			if permuate(board, duplicate, choice, nChar+1, targetStr) {
				return true
			}
			duplicate[choice[0]][choice[1]] = false
		}
		return false
	}
	nRow := len(board)
	nCol := len(board[0])

	initPoints := [][]int{}
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if board[i][j] == byte(word[0]) {
				initPoints = append(initPoints, []int{i, j})
			}
		}
	}

	for _, point := range initPoints {
		duplicate := make([][]bool, nRow)
		for i := 0; i < nRow; i++ {
			duplicate[i] = make([]bool, nCol)
		}
		for i := 0; i < nRow; i++ {
			for j := 0; j < nCol; j++ {
				duplicate[i][j] = false
			}
		}
		duplicate[point[0]][point[1]] = true
		if permuate(board, duplicate, point, 0, word) {
			return true
		}
		duplicate[point[0]][point[1]] = true
	}
	return false
}
