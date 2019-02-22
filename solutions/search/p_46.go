package search

import "log"

type SolutionsFor46 interface {
	Permute(nums []int) [][]int
}

type SolutionsFor46At0 struct{}

// Permute 的这个版本recursive需要花很多时间比较有没有 不好
func (s *SolutionsFor46At0) Permute(nums []int) [][]int {
	combs := [][]int{}

	nNums := len(nums)
	log.Println("choice space:", nums)

	if nNums == 0 {
		return combs
	}

	if nNums == 1 {
		return [][]int{[]int{nums[0]}}
	}

	for index, num := range nums {
		log.Println("choose:", num)
		subCombs := s.Permute(append(nums[:index:index], nums[index+1:]...))
		log.Println("subCombs:", subCombs)
		for _, subComb := range subCombs {
			log.Println("subComb", subComb, "choose:", num)
			newComb := append([]int{num}, subComb...)
			log.Println("subIndex", "newComb", newComb)
			// 如果是组合呢? 就简单的再把位置遍历一下
			combs = append(combs, newComb)
		}
		log.Println("==>combs:", combs)

	}

	return combs
}
