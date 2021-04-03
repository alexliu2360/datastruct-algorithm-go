package queue_stack

// 本题关键是nums1是nums2的子集
// nums1和nums2中所有整数 互不相同
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	var ret []int
	dict := map[int]int{}
	var stack []int
	for _, n := range nums2 {
		//栈顶元素小于当前值，则从栈中取出放入dict中
		//如果栈顶元素大于当前值，则将下一个元素入栈
		for len(stack) > 0 && stack[len(stack)-1] < n {
			dict[stack[len(stack)-1]] = n
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, n)
	}

	//将栈中剩余的值都对应于-1
	for len(stack) > 0 {
		dict[stack[len(stack) - 1]] = -1
		stack = stack[:len(stack) - 1]
	}

	//从dict中取数据放入返回值中
	for _, n := range nums1{
		v := dict[n]
		ret = append(ret, v)
	}
	return ret
}
