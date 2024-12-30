package string

import "fmt"

/*
3.
*/

const LinkNum = 3

type Stringlink struct {
	data []rune
	next *Stringlink
}

func initStringLink(head *Stringlink, str string) *Stringlink {
	var length int
	//i = len(str)
	length = len(str)
	num := length / LinkNum
	if length%LinkNum != 0 {
		num++
	}
	middel := head
	temp := middel
	for i := 0; i < num; i++ {
		if i*LinkNum+LinkNum < length {
			limitStr := str[i*LinkNum : i*LinkNum+LinkNum]
			fmt.Println("limitStr:", limitStr)
			rData := []rune(limitStr)
			temp.data = rData
			var newLink = &Stringlink{}
			temp.next = newLink
			temp = newLink
		} else {
			limitStr := str[i*LinkNum : len(str)]
			fmt.Println("limitStr:", limitStr)
			rData := []rune(limitStr)
			temp.data = rData
			newLink := &Stringlink{}
			temp.next = newLink
			temp = newLink
		}
	}
	return middel
}

func disPlayLink(head *Stringlink) {
	temp := head
	for temp.next != nil {
		//for i := 0; i < LinkNum; i++ {
		fmt.Println(RuneString(temp.data), "->")
		//}
		temp = temp.next
	}
}

func RuneString(value []rune) string {
	return string(value)
}

func StringLinkTest() {
	head := &Stringlink{}
	link := initStringLink(head, "hello, world!")
	disPlayLink(link)
}
