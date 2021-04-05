package design

import (
	"unicode"
)

// stack
// 输入符号  + - * / () " "
// 输入数字  0~9
func isDigit(s rune) (int, bool) {
	switch s {
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		return int(s -'0'), true
	default:
		return 0, false
	}
}

func calculate(s string) int {
	return subCalculate(&s)
}

func subCalculate(s *string) int {
	var (
		num  int    = 0
		sign string = "+"
		res  int    = 0
	)
	stack := make([]int, len(*s))
	for ; len(*s) > 0 ; {
		// 取出第一个字符
		c := (*s)[0]
		*s = (*s)[1:]

		n, ok := isDigit(rune(c))
		if ok {
			num = 10*num + n
		}

		if string(c) == "(" {
			// 递归
			num = subCalculate(s)
		}

		if len(*s) == 0 || (string(c) != " " && !ok) {
			switch sign {
			case "+":
				stack = append(stack, num)
				break
			case "-":
				stack = append(stack, -num)
				break
			case "*":
				// pop
				preNum := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				// push
				stack = append(stack, preNum*num)
				break
			case "/":
				// pop
				preNum := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				// push
				stack = append(stack, preNum/num)
				break
			}
			sign = string(c)
			num = 0
		}
		if string(c) == ")" {
			break
		}
	}

	for len(stack) != 0 {
		res += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}


func calculate2(s string) int {
	stack := make([]int, len(s))
	sign := 1 // 1 为正，-1为负
	res := 0 // res为当前层的计算结果
	num := 0 // num保存一个一位或多位数
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			// 构建多位数（如果有）比如 "12" = 1*10+2 = 12
			j := i // j指针从当前的i开始
			num = 0
			for j < len(s) && unicode.IsDigit(rune(s[j])) { // 如果当前j字符是数字
				num = num*10 + int(s[j]-'0') //计算最新的num
				j++ // j步进
			}
			res += sign * num // 累加最新得出来的结果
			i = j - 1 // i要跟上j，但上面循环中j最后多加了一次1，
		case '(':
			stack = append(stack, res) // 将最新的计算结果res入栈
			stack = append(stack, sign) // 最新的正负符号也入栈，跟随res
			res = 0 // 重置res，因为要准备计算新的层级
			sign = 1 // 重置sign
		case ')':
			sign = stack[len(stack)-1] // 获取栈顶的正负符号
			prevRes := stack[len(stack)-2] // 获取栈顶的上层(之前的)的计算结果
			stack = stack[:len(stack)-2] // 让栈顶的两个元素出栈
			res = prevRes + sign*res // 更新出最新的结果
		case '+':
			sign = 1
		case '-':
			sign = -1
		}
	}
	return res
}