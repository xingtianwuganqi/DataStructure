class Node():
	def __init__(self,item):
		self.element = item
		self.next = None


# 创建单链表类
class SingleLinkList():
	def __init__(self):
		self.header = None # 头指针
		self.length = 0

	# 1.判断是否为空
	def isEmpty(self):
		if self.header == None:
			return True
		else:
			return False


	# 2.头部插入
	def add(self,node):
		if self.isEmpty():
			self.header = node
		else:
			node.next = self.header
			self.header = node
		self.length += 1


	# 3.尾部插入
	def append(self,node):
		currentNode = self.header
		if self.isEmpty():
			self.add(node)
		else:
			while (currentNode.next != None):
				currentNode = currentNode.next # 找到最后的节点
			currentNode.next = node
		self.length += 1


	# 4.指定位置插入
	def insert(self,node,index):
		currentNode = self.header
		if index > self.length + 1 or index <= 0:
			print('插入位置不对，请重新选择位置')
		
		if index == 1:
			self.add(node)
			self.length += 1
		elif index == 2:
			node.next = self.header.next
			self.header.next = node
			self.length += 1
		else:
			for i in range(1,index-1):
				currentNode = currentNode.next
			node.next = currentNode.next
			currentNode.next = node
			self.length += 1

	# 5.遍历
	def travel(self):
		currentNode = self.header
		if self.length == 0:
			print('要遍历的链表没有数据')
			return 
		else:
			for i in range(self.length):
				print('%s'%currentNode.element,end=" ")
				currentNode = currentNode.next
			print('\n')
	# 6.排序不用交换节点的位置，只需要交换节点上的数据值
	def list_sort(self):
		for i  in range(0,self.length-1):
			currentNode = self.header
			for j in range(0,self.length - i - 1):
				if currentNode.element > currentNode.next.element:
					temp = currentNode.element
					currentNode.element = currentNode.next.element
					currentNode.next.element = temp
				currentNode = currentNode.next


	# 7、按索引删除
	def remove(self,index):
		if index < 0 or index > self.length:
			print('输入的下标不对，请重新输入')
			return 
		else:
			if index == 1:
				self.header = self.header.next
				currentNode = self.header
			elif index == 2:
				currentNode = self.header
				currentNode.next = currentNode.next.next
			else:
				currentNode = self.header
				for i in range(1,index-1):
					currentNode = currentNode.next
				currentNode.next = currentNode.next.next
		self.length -= 1






