package main

func isValid(s string) bool {

	var charList []rune

	for _, char := range s {

		// 순서대로 열림 문자일때는 push
		// 닫힘 문자 일때는 push 목록에서 찾아서 제거.

		if char == '(' || char == '{' || char == '[' {

			charList = append([]rune{char}, charList...)

		} else if len(charList) > 0 && (char == ')' || char == '}' || char == ']') {

			if (char == ')' && charList[0] == '(') ||
				(char == '}' && charList[0] == '{') ||
				(char == ']' && charList[0] == '[') {
				charList = charList[1:]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return len(charList) == 0

}
