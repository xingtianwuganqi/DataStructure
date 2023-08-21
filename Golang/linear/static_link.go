package linear

import "fmt"

//import "fmt"

/*
对静态链表的解释：
首先我们让数组的元素都是由两个数据域组成,data和cur。也就是说,数组的每个元素都对应-个data和-个cur。数据域data’用来存放数据元素,也就是通常我们要处理的数据;而cur相当干单链表中的next指针,存放该元素的后继在数组中的下标’我们把cur叫做游标°
我们把这种用数组描述的链表叫做静态链表’这种描述方法还有起名叫做游标实
现法°
Ntote：
数组中未存储数据的元素，称为备用链表
数组的第0个元素和最后一个元素不存储数据
数组中第0个元素的cur存储备用链表的第一个元素的下标
数组的最后一个元素的cur存储第一个拥有真实数据元素的下标

原文链接：https://blog.csdn.net/baidu_41497932/article/details/128420382
*/

const maxSize = 10

type StaticNode struct {
	data int
	cur  int
}

type StaticList struct {
	Element [maxSize]StaticNode
	length  int
}

// InitList 链表初始化
func (link *StaticList) InitList() {
	link.length = 0
	for i := 0; i < maxSize-1; i++ {
		link.Element[i].cur = i + 1
	}
	//目前静态链表为空，最后一个元素的下标为0
	link.Element[maxSize-1].cur = 0
}

// MallocList 返回备用链表的下标，表满返回0
func (link *StaticList) MallocList() int {
	//数组的第一个元素的cur存储备用链表的第一个元素的下标，表满的space[0].cur=0
	i := link.Element[0].cur
	if link.Element[0].cur != 0 {
		link.Element[0].cur = link.Element[i].cur
	}
	return i
}

// FreeList 将下标为k的节点回收到备用链表
func (link *StaticList) FreeList(k int) {
	link.Element[k].cur = link.Element[0].cur
	link.Element[0].cur = k
}

// SSLLength 获取静态链表的长度
func (link *StaticList) SSLLength() int {
	j := 0
	//获取链表中第一个元素的下标（在数组中的最后一个元素存储），数组的最后一个元素的cur存储第一个拥有真实数据元素的下标
	i := link.Element[maxSize-1].cur
	//链表中的末尾元素的cur=0
	for i != 0 {
		i = link.Element[i].cur
		j += 1
	}
	return j
}

// 输出

// 显示链表结构
func (link *StaticList) traverse() {
	for _, v := range link.Element {
		fmt.Printf("%5d", v.cur)
	}
	fmt.Println()
	for _, v := range link.Element {
		fmt.Printf("%5d", v.data)
	}
	fmt.Println()
	for i, _ := range link.Element {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
}

// 判断是否为空
func (link *StaticList) ListEmpty() bool {
	if link.length > 0 {
		return false
	}
	return true
}

// 插入元素
func (link *StaticList) listInsert(value, index int) bool {

	if index < 1 || index > link.SSLLength()+1 {
		fmt.Println("插入位置错误")
		return false
	}
	// 获取备用链的位置，表满返回0
	j := link.MallocList()
	fmt.Println("....,", j)
	// k首先是最后一个元素的下标
	k := maxSize - 1
	if j != 0 {
		link.Element[j].data = value
		for m := 1; m < index; m++ {
			k = link.Element[k].cur
		}
		//循环完k是第i-1个元素的下标,把i-1的cur赋值给L[j].cur
		link.Element[j].cur = link.Element[k].cur
		link.Element[k].cur = j
		return true
	}
	return false
}

// MARK: 删除元素
//func (link *StaticList) DeleteLink(index int, node *StaticNode) {
//	if index < 1 || index > link.SSLLength() {
//		fmt.Println("插入位置错误")
//		return
//	}
//
//	// 判断表非空
//	k := maxSize - 1
//	// 找到第i-1个元素的数组下标
//	for l := 1; l < index; l++ {
//		k = link.Element[k].cur
//	}
//	// j为要删除元素的下标
//	j := link.Element[k].cur
//	*node = link.Element[j].data
//	link.Element[k].cur = link.Element[j].cur
//	link.FreeList(index)
//	return
//}

func StaticListTest() {
	link := new(StaticList)
	link.InitList()
	link.traverse()
	isEmpty := link.ListEmpty()
	fmt.Println(isEmpty)
	link.listInsert(2, 1)
	link.listInsert(3, 2)
	link.traverse()
	length := link.SSLLength()
	fmt.Println(length)
}
