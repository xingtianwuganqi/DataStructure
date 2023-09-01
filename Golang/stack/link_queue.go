package stack

import "fmt"

type qNode struct {
	data int
	next *qNode
}

func initLinkQueue() *qNode {
	node := &qNode{data: 0, next: nil}
	return node
}

// 链式队列入队
func enLinkQueue(rear *qNode, data int) *qNode {
	//1、用节点包裹入队元素
	enElem := &qNode{data: data, next: nil}
	//2、新节点与rear节点建立逻辑关系
	rear.next = enElem
	//3、rear指向新节点
	rear = enElem
	//返回新的rear，为后续新元素入队做准备
	return rear
}

// 链式队列出队
func deLinkQueue(top *qNode, rear *qNode) *qNode {
	var p *qNode
	if top.next == nil {
		fmt.Println("队列为空")
		return rear
	}
	// 1、创建新指针 p 指向目标结点
	p = top.next
	//2、将目标结点从链表上摘除
	top.next = p.next
	if rear == p {
		rear = top
	}
	return rear
}

func printLinkQueue(head *qNode) {
	var temp *qNode
	temp = head
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
}

func LinkQueueTest() {
	var queue, top, rear *qNode
	rear = initLinkQueue()
	top = rear
	queue = top

	rear = enLinkQueue(rear, 1)
	rear = enLinkQueue(rear, 2)
	rear = enLinkQueue(rear, 3)
	rear = enLinkQueue(rear, 4)
	printLinkQueue(queue)
	// 出队
	rear = deLinkQueue(top, rear)
	rear = deLinkQueue(top, rear)
	rear = deLinkQueue(top, rear)
	rear = deLinkQueue(top, rear)
	fmt.Println("\n")
	printLinkQueue(queue)
}
