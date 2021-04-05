package queue_stack

import (
	"fmt"
	"testing"
)

func TestNextGreaterElement1(t *testing.T) {
	nums1 := []int{4,1,2}
	nums2 := []int{1,3,4,2}

	n := nextGreaterElement1(nums1, nums2)
	fmt.Println(n)
}

func TestNextGreaterElement2(t *testing.T) {
	nums1 := []int{1,2,1}
	//nums2 := []int{1,3,4,2}

	n := nextGreaterElement2(nums1)
	fmt.Println(n)
}

func TestDailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	//nums2 := []int{1,3,4,2}

	n := dailyTemperatures(temperatures)
	fmt.Println(n)
}
func TestImplementQueueUsingStacks(t *testing.T) {
	obj := Constructor()
	obj.Push(11)
	obj.Push(22)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
}

func TestImplementStackUsingQueues(t *testing.T) {
	obj := ConstructorMyStack()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}


func TestMaxSlidingWindow(t *testing.T) {
	arr := []int{1,3,-1,-3,5,3,6,7}

	res := maxSlidingWindow(arr, 3)
	fmt.Println(res)
	//1617350108 1617350119
	//fmt.Println(time.Now().Unix())
}

func TestBracket(t *testing.T){
	fmt.Println("1: ", isValid("([{}])"))
	fmt.Println("2: ", isValid("([[{}]])"))
	fmt.Println("3: ", isValid("([{{])"))
	fmt.Println("4: ", isValid("([{}]]]])"))
	fmt.Println("5: ", isValid("([[[[{}])"))
	fmt.Println("6: ", isValid("["))
	fmt.Println("7: ", isValid("]"))
}

func TestBracketisValidWithReplace(t *testing.T){
	fmt.Println("1: ", isValidWithReplace("([{}])"))
	fmt.Println("2: ", isValidWithReplace("([[{}]])"))
	fmt.Println("3: ", isValidWithReplace("([{{])"))
	fmt.Println("4: ", isValidWithReplace("([{}]]]])"))
	fmt.Println("5: ", isValidWithReplace("([[[[{}])"))
	fmt.Println("6: ", isValidWithReplace("["))
	fmt.Println("7: ", isValidWithReplace("]"))
}
func TestBracketTwoChars(t *testing.T){
	fmt.Println("1: ", isValidTwoChars("([{}])"))
	fmt.Println("2: ", isValidTwoChars("([[{}]])"))
	fmt.Println("3: ", isValidTwoChars("([{{])"))
	fmt.Println("4: ", isValidTwoChars("([{}]]]])"))
	fmt.Println("5: ", isValidTwoChars("([[[[{}])"))
	fmt.Println("6: ", isValidTwoChars("["))
	fmt.Println("7: ", isValidTwoChars("]"))
}


























