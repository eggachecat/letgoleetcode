package search

type SolutionsFor70 interface {
	ClimbStairs(n int) int
}

type SolutionsFor70At0 struct{}

// f(k)是走到k
// 所以先到f(k-1)走1 和 f(k-2)走2
// 斐波那契额数列
func (s *SolutionsFor70At0) ClimbStairs(n int) int {
	step := 1
	prev := 0
	for i := 0; i < n; i++ {
		step = prev + step
		prev = step - prev
	}
	return step
}
