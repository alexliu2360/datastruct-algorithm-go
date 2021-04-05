package math


//两个乘积就是前面两个反过来，反转临界点就在 `sqrt(n)`
func isPrim(n int) bool {
	if n == 1{
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
