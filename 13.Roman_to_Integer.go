package main

/*

IV,IX 4,9
XL,XC 40,90
CD,CM 400,900
I = 1
V = 5
X = 10
L = 50
C = 100
D = 500
M = 1000

*/

func romanToInt(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {

		if s[i] == 'C' {
			total += 100
			if i+1 == len(s) {
				break
			} else if s[i+1] == 'D' {
				i++
				total += 300
			} else if s[i+1] == 'M' {
				i++
				total += 800
			}
		} else if s[i] == 'X' {
			total += 10
			if i+1 == len(s) {
				break
			} else if s[i+1] == 'L' {
				i++
				total += 30
			} else if s[i+1] == 'C' {
				i++
				total += 80
			}
		} else if s[i] == 'I' {
			total += 1
			if i+1 == len(s) {
				break
			} else if s[i+1] == 'V' {
				i++
				total += 3
			} else if s[i+1] == 'X' {
				i++
				total += 8
			}
		} else if s[i] == 'V' {
			total += 5
		} else if s[i] == 'L' {
			total += 50
		} else if s[i] == 'D' {
			total += 500
		} else if s[i] == 'M' {
			total += 1000
		}

	}

	return total
}
