package linear

import "fmt"

// ListNode 链表节点
type ListNode struct {
	data int
	next *ListNode
}

// LinkedList  定义链表结构
type LinkedList struct {
	head *ListNode
}

// CreateLinkedList 创建链表
func CreateLinkedList() *LinkedList {
	// 首先定义一个头指针
	return &LinkedList{
		head: nil,
	}
}

// AddNodes 添加几个节点
func (link *LinkedList) AddNodes(num int) {
	for i := 0; i < num; i++ {
		link.Insert(i)
	}
}

// Insert 添加元素
func (link *LinkedList) Insert(value int) {
	node := &ListNode{data: value, next: nil}
	if link.head == nil {
		link.head = node
		return
	}
	current := link.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

// InsertIndex 插入节点
func (link *LinkedList) InsertIndex(value, index int) {
	if index == 1 {
		node := &ListNode{data: value, next: nil}
		next := link.head.next
		node.next = next
		link.head.next = node
	} else {
		current := link.head
		for i := 1; i < index; i++ {
			current = current.next
			if current.next == nil {
				fmt.Println("插入位置无效")
				return
			}
		}
		c := &ListNode{data: value, next: nil}
		c.next = current.next
		current.next = c
	}
}

// DeleteNode 删除节点
func (link *LinkedList) DeleteNode(index int) {
	current := link.head
	for i := 1; i < index; i++ {
		current = current.next
		if current.next == nil {
			fmt.Println("插入位置无效")
			return
		}
		// 找到目标元素的前驱节点
	}
}

// ChangeNode 改节点
func (link *LinkedList) ChangeNode(value, index int) {

}

// Print 打印元素
func (link *LinkedList) Print() {
	current := link.head
	for current != nil {
		fmt.Print(current.data, "->")
		current = current.next
	}
	fmt.Println("nil")
}

func LinkTest() {
	link := CreateLinkedList()
	link.AddNodes(5)
	link.Print()
	link.InsertIndex(5, 1)
	link.Print()

	link.InsertIndex(5, 3)
	link.Print()
}
