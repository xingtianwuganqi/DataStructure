
class Sequence_Table():
	# 初始化
	def __init__(self):
		self.data = [None] * 10
		self.length = 0

	def isFull(self):
		if self.length > 10:
			print('该顺序表已满，无法添加元素')
			return True
		else:
			return False

	# 下标索引查找
	def selectByIndex(self,index):
		if index >= 0 and index <= self.length - 1:
			return self.data[index]
		else:
			print('输入的下标不正确')
			return None

	# 元素查找下标
	def selectByNum(self,num):
		isContain = 0
		for i in range(0,self.length):
			if self.data[i] == num:
				isContain = i
				return i
			


	# 追加数据
	def addNum(self,num):
		if self.isFull() == False:
			self.data[self.length] = num
			self.length += 1

	# 打印顺序表
	def printAllNum(self):
		for i in range(self.length):
			print('a[%s]=%s'%(i,self.data[i]),end=" ")
		print('\n')

	# 按下标插入数据
	def insertNumByIndex(self,num,index):
		if index < 0 or index > self.length:
			return 
		self.length += 1
		# print('self.length-1',self.length-1)
		# print('index:',index)
		for i in range(self.length - 1,index, -1):
			# print(i)
			temp = self.data[i] # 先把None拿出来
			# print('self.data[i-1]',self.data[i-1])
			self.data[i] = self.data[i-1]
			self.data[i-1] = temp # 再把None放到移开的地方
			# print('temp',temp)
		self.data[index] = num


	# 按下标删除
	def delectNumByIndex(self,index):
		if self.length < 0:
			return
		for i in range(index,self.length-1):
			print('i:',i)
			temp = self.data[i]
			print(temp)
			self.data[i] = self.data[i+1]
			self.data[i + 1] = temp
			self.printAllNum
		self.data[self.length-1] = None
		self.length -= 1




if __name__ == "__main__":
	seq_t = Sequence_Table()
	seq_t.addNum(1)
	seq_t.addNum(2)
	seq_t.addNum(3)

	seq_t.printAllNum()

	# 按索引查找
	num = seq_t.selectByIndex(2)

	# 按数字查下标
	index = seq_t.selectByNum(2)

	# 按下标插入
	seq_t.insertNumByIndex(4,2)
	seq_t.printAllNum()


	# 按下标删除
	seq_t.delectNumByIndex(2)

	seq_t.printAllNum()









