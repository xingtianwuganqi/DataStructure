package linear

import "fmt"

type DoubleLinkNode struct {
	prior *DoubleLinkNode
	data  int
	next  *DoubleLinkNode
}

func initDoubleLink(head *DoubleLinkNode) *DoubleLinkNode {
	var list *DoubleLinkNode
	list = head
	for i := 2; i <= 5; i++ {
		// 创建并初始化一个新结点
		node := &DoubleLinkNode{next: nil, prior: nil, data: i}
		// 直接前驱结点的next指向新结点
		list.next = node
		// 新结点指向直接前驱结点
		node.prior = list
		// 下次创建的结点再指向list结点
		list = list.next
	}
	return head
}

// 输出链表数据
func display(head *DoubleLinkNode) {
	var temp = head
	for temp != nil {
		if temp.next == nil {
			fmt.Printf("%d\n", temp.data)
		} else {
			fmt.Printf("%d<->", temp.data)
		}
		temp = temp.next
	}
}

// 插入
func insertLine(head *DoubleLinkNode, data int, add int) *DoubleLinkNode {
	// 新建结点
	var temp = &DoubleLinkNode{prior: nil, data: data, next: nil}
	// 插入到链表头，要特殊考虑
	if add == 1 {
		temp.next = head
		head.prior = temp
		head = temp
	} else {
		var body = head
		// 找到要插入位置的前一个结点
		for i := 1; i < add-1; i++ {
			body = body.next
			// 只要body不存在，表明插入的位置输入错误
			if body == nil {
				fmt.Println("插入位置有误\n")
				return head
			}
		}
		// 判断条件为真，说明插入位置为链表尾
		if body != nil && body.next == nil {
			body.next = temp
			temp.prior = body
		} else {
			body.next.prior = temp
			temp.next = body.next
			body.next = temp
			temp.prior = body
		}
	}
	return head
}

//删除结点的函数，data为要删除结点的数据域的值
func deleteData(head *DoubleLinkNode, data int) *DoubleLinkNode {
	var temp = head
	for temp != nil {
		if temp.data == data {
			// 删除表头结点
			if temp.prior == nil {
				head = head.next
				if head != nil {
					head.prior = nil
					temp.next = nil
				}
				return head
			}
			// 删除表中结点
			if temp.prior != nil && temp.next != nil {
				temp.next.prior = temp.prior
				temp.prior.next = temp.next
				return head
			}
			// 删除表末尾结点
			if temp.next == nil {
				temp.prior.next = nil
				temp.prior = nil
				return head
			}
		}
		temp = temp.next
	}
	return head
}

//head为原双链表，elem表示被查找元素
func selectElem(head *DoubleLinkNode, elem int) int {
	//新建一个指针t，初始化为头指针 head
	var temp = head
	i := 1
	for temp != nil {
		if temp.data == elem {
			return i
		}
		i++
		temp = temp.next
	}
	//程序执行至此处，表示查找失败
	return -1
}

//更新函数，其中，add 表示要修改的元素，newElem 为新数据的值
func amendElem(head *DoubleLinkNode, oldElem int, newElem int) {
	var temp = head
	var find = 0
	for temp != nil {
		if temp.data == oldElem {
			find = 1
			break
		}
		temp = temp.next
	}
	if find == 1 {
		temp.data = newElem
		return
	}
	fmt.Println("修改失败")
}

func DoubleLinkedListTest() {
	var head *DoubleLinkNode = &DoubleLinkNode{
		prior: nil,
		data:  1,
		next:  nil,
	}
	var list = initDoubleLink(head)
	display(list)

	insertLine(list, 6, 3)
	display(list)
	deleteData(list, 6)
	display(list)
	num := selectElem(list, 5)
	fmt.Println(num)
	amendElem(list, 3, 6)
	display(list)
}
