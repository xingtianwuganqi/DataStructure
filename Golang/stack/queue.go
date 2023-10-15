package stack

/*
队列
*/
import "fmt"

type SequenceQueue struct {
	list []int
}

/*
实际上，对于顺序队列的常规实现，我们是从尾部入队列，从下标为0的位置出队列，而不是反过来。

原因在于切片的特性以及队列的先进先出（FIFO）性质：

从尾部入队列： 从切片的尾部添加元素是更高效的，因为切片会自动管理底层数组的扩容。切片会根据需要自动调整容量，当切片的容量不足时，会分配一个更大的底层数组，将原有元素复制到新数组中，这会使插入操作更高效。

从下标为0的位置出队列： 从切片的开头删除元素是更高效的，因为切片的元素是连续存储的，删除开头元素只需要将切片的起始指针向后移动一个位置，而不需要复制大量的元素。

所以，顺序队列一般是从尾部入队列，从下标为0的位置出队列，以充分利用切片的特性和提高操作的效率。非常抱歉之前的回答可能造成了混淆。
*/

// 从尾部入队列，从顶部出队列，可以减少元素挪动的次数，减少时间复杂度
// 从顶部入队列，从尾部出队列，入队列时需要挪动每个元素

func (q *SequenceQueue) enQueue(item int) {
	q.list = append(q.list, item)
}

func (q *SequenceQueue) deQueue() {
	if len(q.list) > 0 {
		q.list = q.list[1:]
	} else {
		panic("队列为空")
	}
}

func (q *SequenceQueue) isEmpty() bool {
	if len(q.list) > 0 {
		return false
	} else {
		return true
	}
}

func (q *SequenceQueue) printQueueItem() {
	for i := 0; i < len(q.list); i++ {
		if i == len(q.list)-1 {
			fmt.Printf("%d", q.list[i])
		} else {
			fmt.Printf("%d -> ", q.list[i])
		}
	}
}

func QueueTest() {
	q := &SequenceQueue{}
	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	empty := q.isEmpty()
	fmt.Println(empty)
	q.printQueueItem()
	q.deQueue()
	fmt.Println("\n")
	q.printQueueItem()
	fmt.Println("\n")
	q.deQueue()

	q.printQueueItem()
}
