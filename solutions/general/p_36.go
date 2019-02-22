package general

import (
	"log"
	"strconv"
)

type SolutionsFor36 interface {
	IsValidSudoku(board [][]byte) bool
}

type SolutionsFor36At0 struct{}

func validateSubBoard(boardInt [][]int) bool {
	// 9 * 9 一定都不一样
	duplicateMap := make([]bool, 10)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if boardInt[i][j] == 0 {
				continue
			}
			if duplicateMap[boardInt[i][j]] {
				return false
			}
			duplicateMap[boardInt[i][j]] = true
		}
	}
	return true
}

func splitBoard(boardInt [][]int) [][][]int {
	subBoards := [][][]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			subBoard := [][]int{
				[]int{boardInt[0+3*i][0+3*j], boardInt[0+3*i][1+3*j], boardInt[0+3*i][2+3*j]},
				[]int{boardInt[1+3*i][0+3*j], boardInt[1+3*i][1+3*j], boardInt[1+3*i][2+3*j]},
				[]int{boardInt[2+3*i][0+3*j], boardInt[2+3*i][1+3*j], boardInt[2+3*i][2+3*j]},
			}
			subBoards = append(subBoards, subBoard)
		}
	}
	return subBoards
}

func parseBoardByte(board [][]byte) [][]int {
	boardInt := [][]int{}
	for _, rowBytes := range board {
		row := string(rowBytes)
		rowInt := []int{}
		for _, ele := range row {
			if string(ele) == "." {
				rowInt = append(rowInt, 0)
			} else {
				eleInt, _ := strconv.Atoi(string(ele))
				rowInt = append(rowInt, eleInt)
			}
		}
		boardInt = append(boardInt, rowInt)
	}
	return boardInt
}

func (s *SolutionsFor36At0) IsValidSudoku(board [][]byte) bool {

	boardInt := parseBoardByte(board)
	// 验证大的 9行 9列
	// 0~8表示行 9~17表示列
	duplicateMap := make([][]bool, 20)
	for i := 0; i < 20; i++ {
		duplicateMap[i] = make([]bool, 10)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := boardInt[i][j]
			if num == 0 {
				continue
			}
			if duplicateMap[i][num] || duplicateMap[9+j][num] {
				log.Println(i, j, duplicateMap[i][num], duplicateMap[9+j][num])
				return false
			}
			duplicateMap[i][num] = true
			duplicateMap[9+j][num] = true
		}
	}

	// 验证小的九宫格
	subBoards := splitBoard(boardInt)
	for _, subBoard := range subBoards {
		if !validateSubBoard(subBoard) {
			log.Println(subBoard)
			return false
		}
	}
	return true
}
