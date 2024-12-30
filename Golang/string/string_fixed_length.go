package string

import "fmt"

/*
对于由多个字符（≥ 0）组成的字符串（例如 "http://data.biancheng.net"），数据结构单独提供了一种存储结构，称为串结构。

字符串中的字符之间具有“一对一”的逻辑关系，所以严格意义上讲，串存储结构也属于线性存储结构。和顺序表、链表、栈、队列这些线性存储结构不同的是，串存储结构专门用来存储字符串。

1. 串的定长顺序存储结构
在 Go 语言中，你可以使用数组来实现定长顺序存储结构，用于表示串（字符串）。串是一种字符序列，因此你可以使用字符数组来表示定长的串
*/

const MaxLength = 10 // 定长串的最大长度

type MyString struct {
	Data   [MaxLength]rune // 使用 rune 类型以支持 Unicode 字符
	Length int
}

// NewMyString 创建一个新的定长串
func NewMyString(s string) *MyString {
	if len(s) > MaxLength {
		s = s[:MaxLength] // 如果字符串超过最大长度，截取前 MaxLength 个字符
	}

	runes := []rune(s)

	var data [MaxLength]rune

	for i := 0; i < len(runes); i++ {
		data[i] = runes[i]
	}
	return &MyString{
		Data:   data,
		Length: len(runes),
	}
}

func (str *MyString) Insert(index int, s string) {
	if index < 0 || index > str.Length {
		fmt.Println("index is Error")
		return
	}
	runes := []rune(s)
	spaceLeft := MaxLength - str.Length
	if len(runes) > spaceLeft {
		fmt.Println("not enought space")
		return
	}

	// 移动元素
	for i := str.Length - 1; i >= index; i-- {
		str.Data[i+len(runes)] = str.Data[i]
	}

	// 插入元素
	for i := index; i < index+len(runes); i++ {
		str.Data[i] = runes[i-index]
	}
	str.Length += len(runes)
}

func (str *MyString) String() string {
	return string(str.Data[:str.Length])
}

func StrListTest() {
	var lens = len("Hello,world!")
	fmt.Println(lens)
	myStr := NewMyString("Hello,world!")
	fmt.Println(myStr.String()) // 输出 "Hello, 世界!"

	data := NewMyString("hell")
	data.Insert(3, "world")
	fmt.Println(data.String())
}
