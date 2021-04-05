package queue_stack

import "strings"

func isValid(s string) bool {
	var stack = make([]string, 0)
	for _, bracket := range s {
		//fmt.Println(i, string(bracket))
		strBracket := string(bracket)
		switch strBracket {
		case "(", "[", "{":
			stack = append(stack, strBracket)
		case ")":
			if len(stack) == 0 || stack[len(stack)-1] != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		case "]":
			if len(stack) == 0 || stack[len(stack)-1] != "[" {
				return false
			}
			stack = stack[:len(stack)-1]
		case "}":
			if len(stack) == 0 || stack[len(stack)-1] != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isValidWithReplace(s string) bool {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "[]", "")
	}

	return len(s) == 0
}

func isValidTwoChars(s string) bool {
	var stack = make([]string, 0)
	for _, bracket := range s {
		char := string(bracket)
		if len(stack) == 0 {
			stack = append(stack, char)
		} else if isSym(stack[len(stack)-1], char) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func isSym(s1 string, s2 string) bool {
	return (s1 == "(" && s2 == ")") || (s1 == "[" && s2 == "]") || (s1 == "{" && s2 == "}")
}
