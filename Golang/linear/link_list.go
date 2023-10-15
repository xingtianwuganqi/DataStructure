package linear

import (
	"fmt"
	"strconv"
)

/*
链表

*/

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
func (link *LinkedList) DeleteNode(value int) int {
	find := -1
	current := link.head
	for current != nil {
		// 找到要删除元素的前驱结点
		if current.next.data == value {
			find = 1
			break
		}
		current = current.next
	}
	if find == -1 {
		return find
	} else {
		current.next = current.next.next
		return 1
	}
}

// DeleteIndex 删除节点
func (link *LinkedList) DeleteIndex(index int) int {
	find := -1
	i := 0
	current := link.head
	for current != nil {
		if i+1 == index {
			find = 1
			break
		}
		// 找到删除节点的前驱节点
		current = current.next
		i++
	}
	current.next = current.next.next
	return find
}

// SelectNode 查询
func (link *LinkedList) SelectNode(value int) int {
	var i = 1
	current := link.head
	for current != nil {
		current = current.next
		if current.data == value {
			return i
		}
		i++
	}
	return -1
}

// ChangeNode 更改节点
func (link *LinkedList) ChangeNode(value, index int) int {
	if index < 1 {
		return -1
	}
	i := 1
	current := link.head
	for current != nil {
		current = current.next
		if i == index {
			current.data = value
			return 1
		}
		i++
	}
	return -1
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

// IterationReverse 链表的反转 迭代反转
func (link *LinkedList) IterationReverse() {
	if link.head == nil || link.head.next == nil {
		return
	}

	var beg *ListNode
	mid := link.head
	end := link.head.next
	// 一直遍历
	for true {
		// 只需修改mid的指针的指向
		mid.next = beg
		if end == nil {
			break
		}

		beg = mid
		mid = end
		end = end.next
	}

	link.head = mid
}

// 递归反转
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	newNode := reverseList(head.next)
	head.next.next = head
	head.next = nil
	return newNode
}

// 两个链表相加
func addTwoNumbers(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	l1.Print()
	l2.Print()
	var beg *ListNode
	mid := l1.head
	end := l1.head.next
	for true {
		mid.next = beg
		if end == nil {
			break
		}
		beg = mid
		mid = end
		end = end.next
	}
	l1.head = mid
	l1.Print()

	var beg2 *ListNode
	mid2 := l2.head
	end2 := l2.head.next
	for true {
		mid2.next = beg2
		if end2 == nil {
			break
		}
		beg2 = mid2
		mid2 = end2
		end2 = end2.next
	}
	l2.head = mid2
	l2.Print()

	// 将数字取出来
	current := l1.head
	i := 0
	total1 := ""
	for current != nil {
		if i == 0 && current.data == 0 {
			current = current.next
			i++
			continue
		}
		total1 = total1 + strconv.Itoa(current.data)
		current = current.next
		i++
	}
	fmt.Println("total1")
	fmt.Println(total1)

	current2 := l2.head
	i2 := 0
	total2 := ""
	for current2 != nil {
		if i2 == 0 && current2.data == 0 {
			current2 = current2.next
			i2++
			continue
		}
		total2 = total2 + strconv.Itoa(current2.data)
		current2 = current2.next
		i2++
	}
	fmt.Println("total2")
	fmt.Println(total2)

	total1Num, _ := strconv.Atoi(total1)
	total2Num, _ := strconv.Atoi(total2)
	total := total1Num + total2Num
	totalStr := strconv.Itoa(total)
	fmt.Println(totalStr)
	// 根据数字再生成链表
	link := &LinkedList{head: nil}
	//curData := link.head
	for i := 0; i < len(totalStr); i++ {
		numData, _ := strconv.Atoi(string(totalStr[i]))
		fmt.Println(numData)
		node := &ListNode{data: numData, next: nil}
		//if curData == nil {
		//	curData = node
		//	fmt.Println(curData)
		//	link.head = curData
		//} else {
		//	curData.next = node
		//	fmt.Println(curData)
		//}
		//curData = curData.next
		if link.head == nil {
			link.head = node
		} else {
			link.Insert(numData)
		}
	}
	fmt.Println(link)
	return link
}

func LinkTest() {
	link := CreateLinkedList()
	link.AddNodes(5)
	link.Print()
	link.InsertIndex(5, 1)
	link.Print()

	link.InsertIndex(5, 3)
	link.Print()

	finish := link.DeleteNode(1)
	if finish == 1 {
		fmt.Printf("删除元素%d\n", 1)
	}
	link.Print()
	selValue := link.SelectNode(4)
	fmt.Printf("查询到的元素的下标，%d\n", selValue)
	changeValue := link.ChangeNode(1, 2)
	link.Print()
	fmt.Printf("替换成功？%d\n", changeValue)
	link.DeleteIndex(2)
	link.Print()
	// 迭代法反转
	link.IterationReverse()
	link.Print()
}

// LinkIterationReverseTest 迭代反转法
func LinkIterationReverseTest() {
	l1 := CreateLinkedList()
	l1.AddNodes(4)

	l2 := CreateLinkedList()
	l2.AddNodes(5)

	l3 := addTwoNumbers(l1, l2)
	l3.Print()
}

// 打印
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.data)
		head = head.next
	}
	fmt.Println("nil")
}

// ReverseListTest 反转测试
func ReverseListTest() {
	l1 := CreateLinkedList()
	l1.AddNodes(4)
	l1.Print()

	// 递归法
	node := reverseList(l1.head)
	printList(node)

	// 头插法
	newNode := headReverseList(node)
	printList(newNode)

	printList(node)
	// 就地逆置法
	curNode := currentReverseList(newNode)
	printList(curNode)

	fmt.Println("====")
	fmt.Println(&node)
	fmt.Println(&newNode)
	fmt.Println(&curNode)
}

/*
头插法
*/

func headReverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	var temp *ListNode
	if head == nil || head.next == nil {
		return head
	}

	for head != nil {
		//temp = head  // 取出头结点
		//head = head.next // 原始链表头结点指向原头结点的下一个节点
		//temp.next = newHead // 取出来的头结点指向nil，指向新链表的头结点
		//newHead = temp // 更新新链表的头节点
		temp = head.next    // 保存头结点的下一个节点
		head.next = newHead // 将头结点的next指向nil，指向新链表的头结点
		newHead = head      // 头结点赋值给newHead，更新新链表的头节点
		head = temp         // 移动到原始链表的下一个节点
	}
	return newHead
}

// 就地逆置法
func currentReverseList(head *ListNode) *ListNode {
	var beg *ListNode
	var end *ListNode
	if head == nil || head.next == nil {
		return head
	}

	beg = head
	end = head.next
	for end != nil {
		beg.next = end.next // beg指向end的下一个节点
		end.next = head     //end指向head
		head = end          // head更新为头结点
		end = beg.next      // end指向beg的下一个节点
	}
	return head
}

/*
遍历每个元素，判断是否有相同的元素
时间复杂度 o(n2) n平方
*/

// LinkIntersect 判断两条链表是否相交
func LinkIntersect(l1 *ListNode, l2 *ListNode) bool {
	p1 := l1
	p2 := l2
	for p1 != nil {
		for p2 != nil {
			if p1 == p2 {
				return true
			}
			p2 = p2.next
		}
		p1 = p1.next
	}
	return false
}

/*
如果相交，则交点的最后一个节点是相同的
时间复杂度o(n)
*/

// LinkIntersect1 判断两条链表是否相交
func LinkIntersect1(l1 *ListNode, l2 *ListNode) bool {
	p1 := l1
	p2 := l2
	for p1 != nil {
		p1 = p1.next
	}
	for p2 != nil {
		p2 = p2.next
	}
	if p1 == p2 {
		return true
	}
	return false
}

/*
相交的链表有一段是重合的，只对比重合的一部分就可以
时间复杂度 o(n)
*/

func LinkIntersect2(l1 *ListNode, l2 *ListNode) bool {
	pLong := l1
	pShot := l2
	var temp *ListNode
	var num1, num2, step = 0, 0, 0

	// 获取plong,pshot的长度
	for pLong != nil {
		pLong = pLong.next
		num1 += 1
	}

	for pShot != nil {
		pShot = pShot.next
		num2 += 2
	}
	// 重置plong，和pshot
	pLong = l1
	pShot = l2
	step = num1 - num2
	if num2 > num1 {
		pLong = l2
		pShot = l1
		step = num2 - num1
	}
	//在 plong 链表中找到和 pshort 等长度的子链表
	temp = pLong
	for step > 0 {
		temp = temp.next
		step--
	}

	//逐个比较 temp 和 pshort 链表中的节点是否相同
	for temp != nil && pShot != nil {
		if temp == pShot {
			return true
		}
		temp = temp.next
		pShot = pShot.next
	}
	return false
}
