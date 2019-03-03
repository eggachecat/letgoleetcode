package search

import (
	"log"
	"strconv"
)

type SolutionsFor85 interface {
	MaximalRectangle(matrix [][]byte) int
}

// 愚蠢的方法: 便利之
type SolutionsFor85At0 struct{}
type SolutionsFor85At1 struct{}

type Point struct {
	x int
	y int
}

func (s *SolutionsFor85At0) ToIntMatrix(matrix [][]byte) [][]int {
	nRows := len(matrix)
	nCols := len(matrix[0])
	intMatrix := [][]int{}
	for i := 0; i < nRows; i++ {
		intRow := []int{}
		for j := 0; j < nCols; j++ {
			num, _ := strconv.Atoi(string(matrix[i][j]))
			intRow = append(intRow, num)
		}
		intMatrix = append(intMatrix, intRow)
	}
	return intMatrix
}

func (s *SolutionsFor85At0) IsLeagelRect(matrix [][]int, lt Point, rb Point) bool {
	legel := true
	for i := lt.x; i <= rb.x; i++ {
		for j := lt.y; j <= rb.y; j++ {
			if matrix[i][j] != 1 {
				return false
			}
		}
	}
	return legel
}

func (s *SolutionsFor85At0) MaximalRectangle(matrix [][]byte) int {
	nRows := len(matrix)
	if nRows == 0 {
		return 0
	}
	nCols := len(matrix[0])
	intMatrix := s.ToIntMatrix(matrix)
	area := 0

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			for m := i; m < nRows; m++ {
				for n := j; n < nCols; n++ {
					rb := Point{m, n}
					lt := Point{i, j}
					if s.IsLeagelRect(intMatrix, lt, rb) {
						newArea := (m - i + 1) * (n - j + 1)
						if newArea > area {
							area = newArea
						}

					}
				}
			}
		}
	}
	return area
}

func (s *SolutionsFor85At1) ToIntMatrix(matrix [][]byte) [][]int {
	nRows := len(matrix)
	nCols := len(matrix[0])
	intMatrix := [][]int{}
	for i := 0; i < nRows; i++ {
		intRow := []int{}
		for j := 0; j < nCols; j++ {
			num, _ := strconv.Atoi(string(matrix[i][j]))
			intRow = append(intRow, num)
		}
		intMatrix = append(intMatrix, intRow)
	}
	return intMatrix
}

func (s *SolutionsFor85At1) GetCurrentLeft(l int, row []int) int {

	for {
		if l < 0 || row[l] != 1 {
			return l + 1
		}
		l--
	}
}

func (s *SolutionsFor85At1) GetCurrentRight(r int, row []int) int {
	rowLen := len(row)
	for {
		if r >= rowLen || row[r] != 1 {
			return r
		}
		r++
	}
}

func (s *SolutionsFor85At1) max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func (s *SolutionsFor85At1) min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

// https://leetcode.com/problems/maximal-rectangle/discuss/29054/Share-my-DP-solution
// ["1","0","1","0","0"],
// ["1","0","1","1","1"],
// ["1","1","1","1","1"],
// ["1","0","0","1","0"]
// 策略: 把matrix看成多个直方图, 每一行及其上方的数据都构成一个直方图, 需要考察matrix.size()个直方图
// 对于每个点(row, col), 我们最后都计算以这个点上方的连续的最长的'1'这根柱子往left, right方向延申可以得到的最大的矩形的面积
// 通过这种方法获取的矩形一定会把最大的矩形包含在内
// height[row][col]记录的是(row, col)这个坐标为底座的直方图柱子的高度, 如果这个点是'0', 那么高度当然是0了
// left[row][col]记录的是(row, col)这个坐标点对应的height可以延申到的最左边的位置
// right[row][col]记录的是(row, col)这个坐标点对应的height可以延申到的最右边的位置+1
// 以上面的matrix为例,
// 对于(row=2, col=1)这个点, left=0, right=5, height=1
// 对于(row=2, col=2)这个点, left=2, right=3, height=3
// (2,2)这个点与(2,1)紧挨着,left和right却已经变化如此之大了, 这是因为left和right除了受左右两边的'1'影响, 还受到了其上方连续的'1'的制约
// 由于点(2,2)上有height=3个'1', 这几个'1'的left的最大值作为当前点的left, 这几个'1'的right的最小值作为当前点的right
// 因此, 实际上, 我们是要找以hight对应的这条线段往左右两边移动(只能往全是'1'的地方移动), 可以扫过的最大面积
// 当hight与目标最大矩形区域的最短的height重合时, 最大矩形的面积就找到了, 如上面的例子, 就是点(2,3)或(2,4)对应的height
func (s *SolutionsFor85At1) MaximalRectangle(matrix [][]byte) int {
	intMatrix := s.ToIntMatrix(matrix)
	log.Println(intMatrix)
	nRow := len(matrix)
	nCol := len(matrix[0])
	left := make([][]int, nRow)
	right := make([][]int, nRow)
	height := make([][]int, nRow)
	for i := 0; i < nRow; i++ {
		left[i] = make([]int, nCol)
		right[i] = make([]int, nCol)
		height[i] = make([]int, nCol)
	}
	area := 0
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if intMatrix[i][j] == 0 {
				// 让left/right当两端的原因在于
				// 是0的话不应该影响下面的1的左右
				height[i][j] = 0
				left[i][j] = 0
				right[i][j] = nCol + 1
				log.Println("i=", i, "j=", j, "cl=", left[i][j], "cr=", right[i][j], (right[i][j]-left[i][j])*height[i][j])
				continue
			}

			curLeft := s.GetCurrentLeft(j, intMatrix[i])
			curRight := s.GetCurrentRight(j, intMatrix[i])
			if i == 0 {
				height[i][j] = 1
				left[i][j] = curLeft
				right[i][j] = curRight
			} else {
				height[i][j] = height[i-1][j] + 1
				left[i][j] = s.max(left[i-1][j], curLeft)
				right[i][j] = s.min(right[i-1][j], curRight)
			}

			newArea := (right[i][j] - left[i][j]) * height[i][j]
			if newArea > area {
				area = newArea
			}
			log.Println("i=", i, "j=", j, "cl=", left[i][j], "cr=", right[i][j], (right[i][j]-left[i][j])*height[i][j])
		}
	}

	return area
}
