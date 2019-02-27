package search

import "log"

type SolutionsFor62 interface {
	UniquePaths(m int, n int) int
}

type SolutionsFor62At0 struct{}

// 子问题:
// 		到(m-1,n) or (m,n-1)
//		然后再各一种方法到目的地
func (s *SolutionsFor62At0) UniquePaths(m int, n int) int {

	// (m,n)
	// 		先到(m-1,n)
	// 		或者先到(m,n-1)
	steps := [][]int{}
	for i := 0; i < m; i++ {
		row := []int{1}
		for j := 0; j < n-1; j++ {
			if i == 0 {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		steps = append(steps, row)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			steps[i][j] = steps[i-1][j] + steps[i][j-1]
		}
	}
	log.Println(steps)

	return steps[m-1][n-1]
}
