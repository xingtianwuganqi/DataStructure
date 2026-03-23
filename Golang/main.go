package main

import (
	"fmt"
)

func main() {
	//linear.CreateSeqSliceList()
	//linear.SeqSliceListTest()
	//linear.LinkTest()
	//linear.LinkIterationReverseTest()
	//linear.ReverseListTest()
	//linear.StaticListTest()
	//linear.DoubleLinkedListTest()
	//linear.CircularLinkTest()
	//linear.DoubleCircularTest()

	//stack.StackTest()
	//stack.LinkStackTest()
	//stack.MatingStringTest()
	//stack.QueueTest()
	//stack.LinkQueueTest()

	// string.StrListTest()
	//string.StringSliceTest()
	//string.StringLinkTest()
	fmt.Println(solution([]int{1, 1, 2, 2, 3, 3, 4, 5, 5}))
}

func solution(cards []int) int {
	for i := 0; i < len(cards); i++ {
		var value int = 0
		for j := 0; j < len(cards); j++ {
			if i != j && cards[i] == cards[j] {
				value += 1
			}
		}
		if value == 0 {
			return cards[i]
		} else {
			continue
		}
	}
	return 0
}
