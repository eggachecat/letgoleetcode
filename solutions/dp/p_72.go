package search

import "log"

type SolutionsFor72 interface {
	MinDistance(word1 string, word2 string) int
}

type SolutionsFor72At0 struct{}
type SolutionsFor72At1 struct{}

// 首先有两个str
// 所以至少应该是f(i,j)
// 		感觉f(i,j)->第一个string的1~i和第二个string 的1~j的距离
//			可以通过f(i-1,j-1)然后到f(i,j)
//			可以通过f(i-1,j)然后到f(i,j)
//			可以通过f(i,j-1)然后到f(i,j)
/*
从图上看就是可以从上到下,从左到右,以及对角线直接过来
	从上到下
	从左到右
	对角线
	比较最后字符是否一样,如果不是就+1否则就保持不变

	下图中?可以
		从r->ho然后再加个o(在r上)
		从h->ro然后再加个o(在h上)
		从r->h然后再加个o(分别)


word1
  |
+---+---+---+---+---+---+
|   | h | o | r | s | e |     <- word2
+---+---+---+---+---+---+
| r | 1 | 2 | 2 | 3 | 4 |
+---+---+---+---+---+---+
| o | 2 | ? |   |   |   |
+---+---+---+---+---+---+
| s | 3 |   |   |   |   |
+---+---+---+---+---+---+
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (s *SolutionsFor72At0) MinDistance(word1 string, word2 string) int {
	// 可以完成
	// 样子太丑-> 人为的string加一个''即可

	if word1 == word2 {
		return 0
	}

	if word1 == "" {
		return len(word2)
	}

	if word2 == "" {
		return len(word1)
	}

	len1 := len(word1)
	len2 := len(word2)

	distances := make([][]int, len1)
	for i := 0; i < len1; i++ {
		distances[i] = make([]int, len2)
	}

	// 边界状态

	rowUsed := false
	colUsed := false
	if word1[0] == word2[0] {
		distances[0][0] = 0
		rowUsed = true
		colUsed = true
	} else {
		distances[0][0] = 1
	}
	//
	for i := 1; i < len1; i++ {
		if (word1[i] == word2[0]) && (!rowUsed) {
			distances[i][0] = distances[i-1][0]
			rowUsed = true
		} else {
			distances[i][0] = distances[i-1][0] + 1
		}
	}

	for j := 1; j < len2; j++ {
		if (word2[j] == word1[0]) && (!colUsed) {
			distances[0][j] = distances[0][j-1]
			colUsed = true
		} else {
			distances[0][j] = distances[0][j-1] + 1
		}
	}
	// 开始
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if word1[i] == word2[j] {
				distances[i][j] = distances[i-1][j-1] // 最后一个一样 对角线过来
			} else {
				distances[i][j] = min(distances[i-1][j-1]+1, min(distances[i-1][j]+1, distances[i][j-1]+1))
				// distances[i-1][j-1]+1 先到 (i-1,j-1)然后replace
				// distances[i-1][j]+1 先到 (i-1,j)然后insert
				// distances[i][j-1]+1 先到 (i,j-1)然后insert
			}
			log.Println(distances, "---------")

		}
	}

	return distances[len1-1][len2-1]
}

func (s *SolutionsFor72At1) MinDistance(word1 string, word2 string) int {
	// 可以完成
	// 样子太丑-> 人为的string加一个''即可
	word1 = " " + word1
	word2 = " " + word2
	if word1 == word2 {
		return 0
	}
	len1 := len(word1)
	if len1 == 1 {
		return len(word2)
	}
	len2 := len(word2)
	if len2 == 1 {
		return len(word1)
	}

	distances := make([][]int, len1)
	for i := 0; i < len1; i++ {
		distances[i] = make([]int, len2)
	}

	// 边界状态
	for i := 1; i < len1; i++ {
		distances[i][0] = i
	}
	for j := 1; j < len2; j++ {
		distances[0][j] = j
	}

	// 开始
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if word1[i] == word2[j] {
				distances[i][j] = distances[i-1][j-1]
			} else {
				distances[i][j] = min(distances[i-1][j-1]+1, min(distances[i-1][j]+1, distances[i][j-1]+1))
			}
			log.Println(distances, "---------")

		}
	}

	return distances[len1-1][len2-1]
}
