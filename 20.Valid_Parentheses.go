package main

func popLast(charList []rune) (rune, []rune) {
	if len(charList) == 0 {
		return 0, charList // 빈 슬라이스 처리
	}
	last := charList[len(charList)-1]     // 마지막 요소
	newList := charList[:len(charList)-1] // 마지막 요소 제외한 슬라이스
	return last, newList
}

func isValid(s string) bool {

	var charList []rune
	var last rune

	for _, char := range s {

		// 순서대로 열림 문자일때는 push
		// 닫힘 문자 일때는 push 목록에서 찾아서 제거.

		if len(charList) == 0 && (char == ')' || char == '}' || char == ']') {
			return false
		} else if char == '(' || char == '{' || char == '[' {
			//charList = append([]rune{char}, charList...)
			charList = append(charList, char)
			continue
		}

		last, charList = popLast(charList)

		if (char == ')' && last != '(') || (char == '}' && last != '{') || (char == ']' && last != '[') {
			return false
		}
	}

	return len(charList) == 0

}
