package math

import (
	"fmt"
	"testing"
	"time"
)

func TestIsPrim(t *testing.T){
	fmt.Println("1: ", isPrim(1))
	fmt.Println("2: ", isPrim(2))
	fmt.Println("3: ", isPrim(3))
	fmt.Println("4: ", isPrim(4))
	fmt.Println("5: ", isPrim(5))
	fmt.Println("6: ", isPrim(6))
	fmt.Println("7: ", isPrim(7))
	fmt.Println("8: ", isPrim(8))
	fmt.Println("9: ", isPrim(9))
	fmt.Println("10: ", isPrim(10))
}

func TestCountPrimes(t *testing.T){
	//fmt.Println(countPrimes(1))// 0
	//fmt.Println(countPrimes(2))// 0
	//fmt.Println(countPrimes(10))// 2,3,5,7
	//fmt.Println(countPrimes(20))// 2	3	5	7	11	13	17	19
	//fmt.Println(countPrimes(30))// 2,3,5,7

	t1 := time.Now().UnixNano()
	countPrimes(10000)
	fmt.Println((time.Now().UnixNano() -t1)/1000000, "ms")

	t1 = time.Now().UnixNano()
	countPrimes(5 * 1000000)
	fmt.Println((time.Now().UnixNano() -t1)/1000000, "ms")
}

func TestUglyNum(t *testing.T){
	fmt.Println("1: ", isUgly(1))
	fmt.Println("2: ", isUgly(2))
	fmt.Println("3: ", isUgly(3))
	fmt.Println("4: ", isUgly(4))
	fmt.Println("5: ", isUgly(5))
	fmt.Println("6: ", isUgly(6))
	fmt.Println("7: ", isUgly(7))
	fmt.Println("8: ", isUgly(8))
	fmt.Println("9: ", isUgly(9))
	fmt.Println("10: ", isUgly(10))
	fmt.Println("11: ", isUgly(11))
	fmt.Println("14: ", isUgly(14))
	fmt.Println("22: ", isUgly(22))
	fmt.Println("24: ", isUgly(24))
	fmt.Println("-2147483648: ", isUgly(-2147483648))

}