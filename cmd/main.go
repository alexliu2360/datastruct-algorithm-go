package main

import (
	"bytes"
	"code_brush/basic-data-structure/advanced-tree"
	"fmt"
)

func uniqueMorseRepresentations(words []string) int {
	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	buffer := bytes.Buffer{}
	uniqueTree := advanced_tree.NewBST()

	for _, word := range words {
		buffer.Reset()
		for _, letter := range word {
			buffer.WriteString(morseCode[letter-'a'])
		}
		uniqueTree.InsertNode(buffer.String())
	}

	return uniqueTree.GetSize()
}

func main() {
	//bst := advanced_tree.NewBST()
	//bst1 := binary_search_tree.NewBST()
	//nums := [...]int{5, 3, 8, 6, 4, 2}
	//nums := [...]int{1,2,3,4,5,6}
	//for _, num := range nums {
	//	bst.InsertNode(num)
	//}
	//bst.PreOrder()
	//fmt.Println("===")
	//bst.InOrder()
	//fmt.Println("===")
	//bst.PostOrder()
	//fmt.Println("===")
	//bst.LevelOrder()

	//fmt.Println(bst.Maximum())
	//fmt.Println(bst.Minimum())
	//bst.RemoveMax()
	//bst.DeleteNode(5)
	//bst.InsertNode(5)
	//bst.DeleteNode(1)


	//fmt.Println("pride-and-prejudice")
	//
	//filename, _ := filepath.Abs("./data/pride-and-prejudice.txt")
	//fmt.Println(filename)
	//words1 := utils.ReadFile(filename)
	//fmt.Println(len(words1))
	//fmt.Println("=====bst set =====")
	//startTime := time.Now()
	//set1 := set.NewBstSet()
	//for _, word := range words1 {
	//	set1.Add(word)
	//}
	//fmt.Println(set1.GetSize())
	//fmt.Println(time.Now().Sub(startTime))
	//
	//
	//filename, _ = filepath.Abs("./data/pride-and-prejudice.txt")
	//words2 := utils.ReadFile(filename)
	//fmt.Println(len(words2))
	//
	//fmt.Println("=====linked list set=====")
	//startTime = time.Now()
	//set2 := set.NewLLSet()
	//for _, word := range words2 {
	//	set2.Add(word)
	//}
	//fmt.Println(set2.GetSize())
	//fmt.Println(time.Now().Sub(startTime))

	//n := 10
	//
	//maxHeap := heap.NewMaxHeap()
	//rand.Seed(time.Now().Unix())
	//for i := 0; i < n; i++ {
	//	randNum := rand.Intn(math.MaxInt32)
	//	maxHeap.Add(randNum)
	//}
	//
	//arr := make([]int, n)
	//for i := 0; i < n; i++ {
	//	arr[i] = maxHeap.ExtractMax().(int)
	//}
	//fmt.Println(arr)
	//nums := []int{1, 1, 1, 2, 2, 3,4,4,5,5,5,5,6,6,7,7,7,7,7,7}
	//heap.TopKFrequentElements(nums, 2)



	var a int = 1
	var b int = 3
	var c int


	c = a % b

	fmt.Printf("c is %d\n", c )


}
