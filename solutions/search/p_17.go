package search

type SolutionsFor17 interface {
	LetterCombinations(digits string) []string
}

type SolutionsFor17At0 struct{}

func str2charArr(s string) []string {
	arr := []string{}
	for _, char := range s {
		arr = append(arr, string(char))
	}
	return arr
}

func (s *SolutionsFor17At0) LetterCombinations(digits string) []string {
	num2str := map[string]string{
		"2": "abc", "3": "def", "4": "ghi",
		"5": "jkl", "6": "mno", "7": "pqrs",
		"8": "tuv", "9": "wxyz",
	}
	combs := []string{}

	if len(digits) == 0 {
		return combs
	}

	if len(digits) == 1 {
		return str2charArr(num2str[string(digits[0])])
	}

	subCombs := s.LetterCombinations(digits[1:])

	digitPoss := str2charArr(num2str[string(digits[0])])

	for _, char := range digitPoss {
		for _, subChar := range subCombs {
			combs = append(combs, char+subChar)
		}
	}

	return combs
}
