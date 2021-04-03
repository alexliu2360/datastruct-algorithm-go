package queue_stack

// loop array
// [1,2,1] => [2,-1, 2]
// 本题关键
// 1、需要将数组扩容（即搜索两次）
// 2、根据较小值的索引记录下一个更大值
func nextGreaterElement2(nums1 []int) []int {
	var ret []int
	dict := map[int]int{}
	var stack []int

	nums1Len := len(nums1)
	for i := 0; i < nums1Len*2; i++ {
		// 获取nums1中的索引值
		for len(stack) > 0 && nums1[i%nums1Len] > nums1[stack[len(stack)-1]] {
			// 构建dict
			dict[stack[len(stack)-1]] = nums1[i%nums1Len] // dict中key:索引，value：下一个更大的值
			stack = stack[:len(stack)-1] // pop操作
		}
		if i < nums1Len {
			stack = append(stack, i) // push操作 stack中存储的是上一个更小值的索引
		}
	}

	for _, n := range stack {
		dict[n] = -1
		stack = stack[:len(stack)-1]
	}

	for i := 0; i < len(nums1); i++ {
		value := dict[i]
		ret = append(ret, value)
	}

	return ret
}
