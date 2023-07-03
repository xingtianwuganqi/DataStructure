# 变位词：两个词之间存在组成字母的重新排列关系
'''
法①：检查标记【时间复杂度为O(n2)】

　　思路：检查第一个字符串中的所有字符是不是都在第二个字符串中出现。 
如果能够把每一个字符都“检查标记”一遍，那么这两个字符串就互为变位词。
检查标记一个字符 要用特定值 None 来代替，作为标记。
然而，由于字符串不可变，首先要把第二个字符串转化成一个列表。
第一个字符串中的每一个字符都可以在列表的字符中去检查，
如果找到，就用 None 代替以示标记。
时间复杂度 O(n2)
'''
def anagramSolution1(s1,s2):
	alist = list(s2)
	pos1 = 0
	stillOK = True
	while pos1 < len(s1) and stillOK:
		pos2 = 0
		found = False
		while pos2 < len(alist) and not found:
			if s1[pos1] == alist[pos2]:
				found = True
			else:
				pos2 = pos2 + 1

		if found:
			alist[pos2] = None
		else:
			stillOK = False
		pos1 = pos1 + 1

	return stillOK


'''
法②：排序比较【时间复杂度为O(n2)】

　　思路：尽管 s1 和 s2 并不相同，但若为变位词它们一定包含完全一样的字符，
利用这一特点，我们可以 采用另一种方法。
我们首先从 a 到 z 给每一个字符串按字母顺序进行排序，如果它们是变位词，
那么 我们将得到两个完全一样的字符串。此外，我们可以先将字符串转化为列表，再利用 Python 中内建
的 sort 方法对列表进行排序。下面代码展示了这种方法。
第一眼看上去你可能会认为这个算法的复杂度是 O(n)，
毕竟排序后只需要一个简单的循环去比较 n 个字符。
然而对 Python 内建的 sort 方法的两次使用并非毫无消耗。
事实上，正如我们在后面的章节 中将要看到的，
排序方法的复杂度往往都是 O(n²)或者 O(n㏒n)，所以排序贡献了这个函数主要的循 环操作。
最终，这个算法和排序的复杂度相同。
时间复杂度 O(nlogn)
'''

def anagramSolution2(s1,s2):
	alist1 = list(s1)
	alist2 = list(s2)

	alist1.sort()
	alist2.sort()
	pos = 0
	matches = True
	while pos < len(s1) and matches:
		if alist1[pos] == alist2[pos]:
			pos = pos + 1
		else:
			matches = False

	return matches


'''
法③：计数比较法【时间复杂度O(n)】

解决变位词问题的最后一个方法利用了任何变位词都有相同数量的 a，相同数量的 b，相同数量 的 c 等等。
为判断两个字符串是否为变位词，我们首先计算每一个字符在字符串中出现的次数。
由于共有 26 个可能的字符，我们可以利用有 26 个计数器的列表，每个计数器对应一个字符。
每当我们 看到一个字符，就在相对应的计数器上加一。最终，如果这两个计数器列表相同，则这两个字符串 是变位词。
时间复杂度：O(n)
'''

def anagramSolution3(s1,s2):
	c1 = [0] * 26
	c2 = [0] * 26
	print(c1)
	print(c2)
	for i in range(len(s1)):
		pos = ord(s1[i]) - ord('a')
		c1[pos] = c1[pos] + 1

	for i in range(len(s2)):
		pos = ord(s2[i]) - ord('a')
		c2[pos] = c2[pos] + 1
	print(c1)
	print(c2)
	j = 0
	stillOK = True
	while j < 26 and stillOK:
		if c1[j] == c2[j]:
			j = j + 1
		else:
			stillOK = False
	return stillOK

if __name__ == '__main__':
	value = anagramSolution3('earth','heart')
	print(value)