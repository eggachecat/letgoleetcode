package search

type SolutionsFor127 interface {
	LadderLength(beginWord string, endWord string, wordList []string) int
}

type SolutionsFor127At0 struct{}

var Distance [][]int

func (s *SolutionsFor127At0) transformExists(strA string, strB string) bool {
	diff := 0
	for i, a := range strA {
		if byte(a) != strB[i] {
			diff++
		}
	}
	if diff == 1 {
		return true
	}
	return false
}

func (s *SolutionsFor127At0) getChoices(currentWord string, wordList []string, duplicate []bool) []int {
	choices := []int{}
	for index, word := range wordList {
		if s.transformExists(currentWord, word) && !duplicate[index] {
			choices = append(choices, index)
		}
	}
	return choices
}
func do(a int) {

}

// LadderLength blablabla
// Choices: 某一个word
// Constraints: 和当前的word只能差一个字母 ** 不要重复
// Goal: 变成endWord
// 注意 如果是暴力会有个问题: 重复太多, 应该用bfs
// 一点感悟: 如果 choice 会影响到 后面的选择空间: 暴力
// 关于搜索的一点小建议
// If using DFS to compute the shortest path, the order of visiting nodes should first be topologically sorted.
// If it is not, there may be loops in the DFS forest (recursive tree).
// For BFS, the visiting order is naturally topologically sorted.
// 有循环的图应该用BFS而不是DFS
// 此外这一题比较特殊: BFS 先走到 后面再走就没有意义了
// 比如第一层[A,B] 有路径A->C和B->D->C显然第一条先走C被标记
func (s *SolutionsFor127At0) LadderLength(beginWord string, endWord string, wordList []string) int {

	duplicate := make([]bool, len(wordList))
	currentWords := []string{beginWord}
	ctr := 1
	for {
		ctr++
		choices := []int{}
		for _, currentWord := range currentWords {
			subChoices := s.getChoices(currentWord, wordList, duplicate)
			choices = append(choices, subChoices...)
		}
		currentWords = []string{}
		for _, choice := range choices {
			if wordList[choice] == endWord {
				return ctr
			}
			duplicate[choice] = true
			currentWords = append(currentWords, wordList[choice])
		}
		if len(currentWords) == 0 {
			return 0
		}
	}
}
