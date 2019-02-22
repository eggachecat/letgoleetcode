package search

import (
	"log"
	"strings"
)

type SolutionsFor22 interface {
	GenerateParenthesis(n int) []string
}

type SolutionsFor22At0 struct{}
type SolutionsFor22At1 struct{}

// GenerateParenthesis 有限制的 Recursion
// 相比22多了限制!
// Choice: "(" or ")"
// Constraint:
//			必须满足: 剩余Open的数量 < 剩余Close的数量
// 				=> open == close 的时候只能 open
// Goal: 选择 2*n 个 string
func GenerateParenthesis(nOpen int, nClose int) []string {
	combs := []string{}
	var subCombs []string
	log.Println("choice space:", nOpen, nClose)

	if nOpen == 0 {
		return []string{strings.Repeat(")", nClose)}
	}

	if nOpen < nClose {
		// choiceSpace 是 []string{"(", ")"}

		// 即可以选"(" -> 再open一个
		subCombs = GenerateParenthesis(nOpen-1, nClose)
		log.Println("subCombs: open", subCombs, "choose (")

		for _, subComb := range subCombs {
			combs = append(combs, "("+subComb)
		}
		// 也可以选")" -> close掉以前的
		subCombs = GenerateParenthesis(nOpen, nClose-1)
		for _, subComb := range subCombs {
			combs = append(combs, ")"+subComb)
		}
		log.Println("subCombs: close", subCombs, "choose )")

		nOpen--
		nClose--
	}

	if nOpen == nClose {
		// choiceSpace 是 []string{"("}
		nOpen--
		subCombs = GenerateParenthesis(nOpen, nClose)
		log.Println("subCombs: open", subCombs, "choose (")
		for _, subComb := range subCombs {
			combs = append(combs, "("+subComb)
		}
	}

	return combs
}

func (s *SolutionsFor22At0) GenerateParenthesis(n int) []string {
	return GenerateParenthesis(n, n)
}

// GenerateParenthesis 这个是 sample solution
func (s *SolutionsFor22At1) GenerateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	res := []string{}
	// 对不起 我竟然忘掉了
	// 这边是从open,close=0开始
	var permuate func(s string, open, close int)
	permuate = func(s string, open, close int) {
		if len(s) == 2*n {
			// 全部选好
			res = append(res, s)
			return
		}
		if open < n {
			// 还可以继续open
			permuate(s+"(", open+1, close)
		}
		if close < open {
			// 选择close了
			permuate(s+")", open, close+1)
		}
	}
	permuate("", 0, 0)
	return res
}
