package linear

import "fmt"

//import "fmt"

const maxSize = 10

type StaticNode struct {
	data int
	cur  int
}

type StaticList struct {
	Element [maxSize]StaticNode
	length  int
}

func CreateStaticList(num int) *StaticList {
	var list = new(StaticList)
	for i := 0; i < num-1; i++ {
		node := StaticNode{cur: i + 1}
		list.Element[i] = node
	}
	list.Element[num-1].cur = 0
	return list
}

func StaticListTest() {

	link := CreateStaticList(5)
	fmt.Println(link)
}
