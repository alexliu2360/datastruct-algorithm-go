package math

import "fmt"

// 该方法容易超出时间限制
// 超出限制的原因是isPrim函数的复杂度为O(n2)
func countPrimesOverTime(n int) int {
	count := 0
	for i := 1; i < n; i++ {
		if isPrim(i) {
			fmt.Printf("%d\t", i)
			count++
		}
	}
	fmt.Println()
	return count
}

//寻找小于n的所有素数的个数
//hash表，空间换时间
func countPrimes(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return 0
	}
	// 分配n大小的数组，[0...n-1]
	isPrimList := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrimList[i] = true
	}

	for i := 2; i*i < n; i++ {// 只需要遍历 `[2,sqrt(n)]`, 大于sqrt(n)的数据会被小于sqrt(n)的数据标记完
		if isPrimList[i] {
			for j := i * i; j < n; j += i { // i*i之前的数字都被小于i的其他数字标记过了
				isPrimList[j] = false
			}
		}
	}

	res := 0
	for i := 2; i < n; i++ {
		if isPrimList[i] {
			//fmt.Printf("%d\t", i)
			res++
		}
	}
	//fmt.Println()
	return res
}
