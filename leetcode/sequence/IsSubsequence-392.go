package sequence

//动态规划
//二分查找
//子序列问题

//example:
//s = "abc", t = "**a**h**b**gd**c**", return true.
//s = "axc", t = "ahbgdc", return false.

// 双指针
func isSubsequence(s string, t string) bool {
	i := 0
	j := 0
	for ; i < len(s) && j < len(t); {
		if s[i] == t[j]{
			i++
		}
		j++
	}
	return i == len(s)
}
