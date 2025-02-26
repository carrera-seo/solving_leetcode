package main

func longestCommonPrefix(strs []string) string {
	answer := ""
	for i := 0; i < len(strs[0]); i++ {
		curr_str := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || curr_str != strs[j][i] {
				return answer
			}
		}
		answer += string(curr_str)
	}
	return answer
}
