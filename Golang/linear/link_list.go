package linear

// ListNode 链表节点
type ListNode struct {
	data int
	next *ListNode
}

// LinkedList  定义链表结构
type LinkedList struct {
	head *ListNode
}

func CreateLinkedList() *LinkedList {
	// 首先定义一个头指针
	return &LinkedList{
		head: nil,
	}
}
