package math

func isUgly(n int) bool {

	if n == 1 || n == 0 || n == -1 {
		return true
	}

	for n%5 == 0 {
		n /= 5
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%2 == 0 {
		n /= 2
	}
	if n == 1 {
		return true
	}
	return false

}
