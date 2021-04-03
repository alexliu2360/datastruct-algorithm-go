package queue_stack

// [INPUT]: temperatures = [73, 74, 75, 71, 69, 72, 76, 73]
// [OUTPUT]: temperatures = [1, 1, 4, 2, 1, 1, 0, 0]
// list limit: [1, 30000] temperature: [30, 100]
func dailyTemperatures(T []int) []int {
	dict := make(map[int]int)
	var stack []int
	var res []int

	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			dict[stack[len(stack)-1]] = i - stack[len(stack)-1] // 得到间距
			stack = stack[:len(stack)-1]                        //pop
		}
		stack = append(stack, i)
	}

	for i := 0; i < len(T); i++ {
		v := dict[i]
		res = append(res, v)
	}

	return res
}
