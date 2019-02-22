package search

import (
	"log"
	"strconv"
)

type SolutionsFor37 interface {
	SolveSudoku(board [][]byte)
}

type SolutionsFor37At0 struct{}

func parseBoardByte(board [][]byte) ([][]int, [][]int) {
	boardInt := [][]int{}
	points := [][]int{}
	for rowIndex, rowBytes := range board {
		row := string(rowBytes)
		rowInt := []int{}
		for colIndex, ele := range row {
			if string(ele) == "." {
				rowInt = append(rowInt, 0)
				points = append(points, []int{rowIndex, colIndex})
			} else {
				eleInt, _ := strconv.Atoi(string(ele))
				rowInt = append(rowInt, eleInt)
			}
		}
		boardInt = append(boardInt, rowInt)
	}
	return boardInt, points
}

func tranpose(boardInt [][]int) [][]int {
	boardIntTranpose := [][]int{}
	nRow, nCol := len(boardInt), len(boardInt[0])
	for i := 0; i < nCol; i++ {
		rowTranpose := []int{}
		for j := 0; j < nRow; j++ {
			rowTranpose = append(rowTranpose, boardInt[j][i])
		}
		boardIntTranpose = append(boardIntTranpose, rowTranpose)
	}
	return boardIntTranpose
}

// getChoiceSpace 再多做一点包括九宫格
func getChoiceSpace(boardInt [][]int, point []int) []int {

	choices := []int{}
	rowIndex, colIndex := point[0], point[1]
	rowOffset, colOffset := rowIndex/3, colIndex/3
	// log.Println(rowOffset, colOffset)
	for choice := 1; choice < 10; choice++ {
		avaliable := true
		for i := 0; i < 9; i++ {
			// 在那一col 或者 那个row出现
			if choice == boardInt[i][colIndex] || choice == boardInt[rowIndex][i] {
				avaliable = false
				break
			}
		}
		// 九宫格
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if boardInt[rowOffset*3+i][colOffset*3+j] == choice {
					avaliable = false
					break
				}
			}
		}
		if avaliable {
			choices = append(choices, choice)
		}
	}

	return choices
}

func writeBoardIntToByte(boardInt [][]int, boardByte [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			boardByte[i][j] = string(boardInt[i][j])[0]
		}
		// boardByte[i] = []byte(rowStr)
	}
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

func validateSubBoard(boardInt [][]int) bool {
	// 9 * 9 一定都不一样
	duplicateMap := make([]bool, 10)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if duplicateMap[boardInt[i][j]] {
				return false
			}
			duplicateMap[boardInt[i][j]] = true
		}
	}
	return true
}

func validateBoard(boardInt [][]int) bool {
	subBoards := splitBoard(boardInt)
	for _, subBoard := range subBoards {
		if !validateSubBoard(subBoard) {
			return false
		}
	}
	return true
}

func prettyPrint(boardInt [][]int) {
	for _, row := range boardInt {
		log.Println(row)
	}
}

// SolveSudoku blabla
// 数独 分成几个recursive的问题:
//		1. 9个9宫格
// Choice: 可以选的space(挑剩下)
// Constraint:
//			必须满足: 每一行每一列不重复(再Choice space里的限制)
//          某一个可选:= 这一行可选 交集 这一列可选
// Goal: 行/列的和相同
func (s *SolutionsFor37At0) SolveSudoku(board [][]byte) {
	var permuate func(boardInt [][]int, points [][]int) bool
	permuate = func(boardInt [][]int, points [][]int) bool {
		if len(points) == 0 {
			// 没得选了 不需要验证了, 因为九宫格也考虑在选择里
			return true
		}
		// 对于点来做搜索
		// 选择第一个点 并且 他所有可以的
		point := points[0]
		choices := getChoiceSpace(boardInt, point)
		// log.Println("point", point, "Choices:", choices)
		for _, choice := range choices {
			// log.Println("point", point, "choose:", choice)
			boardInt[point[0]][point[1]] = choice
			if permuate(boardInt, points[1:]) {
				return true
			}
			boardInt[point[0]][point[1]] = 0
		}
		// log.Println("point", point, "No choices remain")
		return false
	}

	boardInt, points := parseBoardByte(board)
	permuate(boardInt, points)
	prettyPrint(boardInt)
	writeBoardIntToByte(boardInt, board)
	log.Println(board)
}
