package linear

import "fmt"

/*
已知 n 个人（分别用编号 1，2，3，…，n 表示）围坐在一张圆桌周围，从编号为 k 的人开始顺时针报数，数到 m 的那个人出列；
他的下一个人又从 1 开始，还是顺时针开始报数，数到 m 的那个人又出列；依次重复下去，直到圆桌上剩余一个人。
*/

type circularNode struct {
	data int
	next *circularNode
}

func initCircularLink(num int) *circularNode {
	var head, temp *circularNode
	head = &circularNode{data: 1, next: nil}
	temp = head
	for i := 2; i < num; i++ {
		current := &circularNode{data: i, next: nil}
		temp.next = current
		temp = temp.next
	}
	temp.next = head
	return head
}

// 打印
func printCirList(head *circularNode) {
	var num = 1
	for head != nil {
		fmt.Printf("%d -> ", head.data)
		if num == 5 {
			return
		}
		head = head.next
		num++
	}
	fmt.Println("nil")
}

// 从编号为k的人开始，数到m的人出列
func findAllKillKey(head *circularNode, k, m int) {
	var p *circularNode
	tail := head
	//找到链表第一个结点的上一个结点，为删除操作做准备
	for tail.next != head {
		tail = tail.next
	}
	p = head
	//找到编号为k的人
	for p.data != k {
		p = p.next
	}
	//从编号为k的人开始，只有符合p->next==p时，说明链表中除了p结点，所有编号都出列了，
	for p.next != p {
		//找到从p报数1开始，报m的人，并且还要知道数m-1的人的位置tail，方便做删除操作。
		for i := 1; i < m; i++ {
			tail = p   // 要出列的前一个人
			p = p.next // 往下循环
		}
		tail.next = p.next
		fmt.Printf("出列人编号：%d\n", p.data)
		p = tail.next
	}
	fmt.Printf("胜出人编号====%d\n", p.data)
}

/*
1) 最直接的实现思想就是：从给定链表的第一个节点开始遍历，每遍历至一个节点，
都将其和所有的前驱节点进行 比对，如果为同一个节点，则表明当前链表中有环；
反之，如果遍历至链表最后一个节点，仍未找到相同的节点， 则证明该链表中无环。

注意，如果一个单链表为有环链表，基于单链表中各节点有且仅有 1 个指针域的特性，则势必该链表是没有尾 结点的（如图 1 所示）。
换句话说，有环链表的遍历过程是无法自行结束的，需要使用 break 语句手动结束遍 历。
时间复杂度o(n2)
*/
// 判断链表上有环

func haveRing(head *circularNode) bool {
	var temp = head
	var add [20]*circularNode
	var length = 0
	for temp != nil {
		//依次将 temp 和 add 数组中记录的已遍历的地址进行比对
		for i := 0; i < length; i++ {
			if temp == add[i] {
				return true
			}
		}
		add[length] = temp
		length++
		temp = temp.next
	}
	return false
}

/*
该算法的实现思想需要借助一个论点，即在一个链表中，如果 2 个指针（假设为 H1 和 H2）都从表头开始遍历 链表，
其中 H1 每次移动 2 个节点的长度（H1 = H1->next->next）,而 H2 每次移动 1 个节点的长度（H2 = H2->next），
如果该链表为有环链表，则 H1、H2 最终必定会相等；反之，如果该链表中无环，则 H1、H2 永远 不会相遇。

有关在有环链表中 H1 和 H2 必定会相遇的结论，假设有环链表中的环包含 n 个节点，则第一次遍历，H1 和 H2 相差 1 个节点；
第二次遍历，H1 和 H2 相差 2 个节点；第三次遍历，H1 和 H2 相差 3 个节点...，
最终 经过多次遍历，H1 和 H2 会相差 n-1 个节点，此时就会在环中重合，此时 H1 和 H2 相等,
时间复杂度o(n)
*/

func haveRing2(head *circularNode) bool {
	var h1 = head.next
	var h2 = head
	for h1 != nil {
		if h1 == h2 {
			return true
		} else {
			h1 = h1.next
			if h1 == nil {
				return false
			} else {
				h1 = h1.next
				h2 = h2.next
			}
		}
	}
	return false
}

func CircularLinkTest() {
	head := initCircularLink(6)
	printCirList(head)
	findAllKillKey(head, 3, 2)
	isHave := haveRing(head)
	fmt.Println(isHave)

	hadRing := haveRing2(head)
	fmt.Println(hadRing)
}
