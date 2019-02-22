package search

type SolutionsFor62 interface {
	UniquePaths(m int, n int) int
}

type SolutionsFor62At0 struct{}

// Permute 的这个版本recursive需要花很多时间比较有没有 不好
func (s *SolutionsFor62At0) UniquePaths(m int, n int) int {
	return 0
}
