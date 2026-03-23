//: A UIKit based Playground for presenting user interface
  
import UIKit
import PlaygroundSupport
/*
 1.顺序表
 */
struct SequentialList<T> {
    private var elements: [T] = []

    // 获取当前顺序表的长度
    var length: Int {
        return elements.count
    }

    // 检查顺序表是否为空
    var isEmpty: Bool {
        return elements.isEmpty
    }

    // 添加元素到顺序表末尾
    mutating func append(_ element: T) {
        elements.append(element)
    }

    // 在指定位置插入元素
    mutating func insert(_ element: T, at index: Int) throws {
        guard index >= 0 && index <= elements.count else {
            throw SequentialListError.indexOutOfBounds
        }
        elements.insert(element, at: index)
    }

    // 删除指定位置的元素
    mutating func remove(at index: Int) throws -> T {
        guard index >= 0 && index < elements.count else {
            throw SequentialListError.indexOutOfBounds
        }
        return elements.remove(at: index)
    }

    // 获取指定位置的元素
    func get(at index: Int) throws -> T {
        guard index >= 0 && index < elements.count else {
            throw SequentialListError.indexOutOfBounds
        }
        return elements[index]
    }

    // 修改指定位置的元素
    mutating func update(at index: Int, with element: T) throws {
        guard index >= 0 && index < elements.count else {
            throw SequentialListError.indexOutOfBounds
        }
        elements[index] = element
    }
    
    // 遍历
    func traversal() {
        for i in elements {
            print(i)
        }
    }

    // 清空顺序表
    mutating func clear() {
        elements.removeAll()
    }
    
    // 打印顺序表内容
    func printList() {
        print(elements)
    }

    // 自定义错误类型
    enum SequentialListError: Error {
        case indexOutOfBounds
    }
}

/*
 顺序存储结构优缺点
 优点：
    无须为表中数据元素的逻辑关系增加额外的存储空间。
    可以快速的存取表中任一位置的元素。（时间复杂度为 O(1)）
 缺点:
    插入和删除操作需要移动大量元素。（时间复杂度为 O(n)）
    当线性表长度变化较大时，难以确定存储空间的容量。
    造成存储空间的碎片。
 */
do {
    var list = SequentialList<Int>()
    list.append(10)
    list.append(20)
    list.append(30)
    try list.insert(15, at: 1) // 在索引 1 处插入 15
    print("List after insertion:")
    list.printList() // 打印 [10, 15, 20, 30]
    print("遍历：")
    list.traversal()

    let removedElement = try list.remove(at: 2) // 删除索引 2 处的元素
    print("Removed element:", removedElement) // 打印 20
    list.printList() // 打印 [10, 15, 30]

    let element = try list.get(at: 1) // 获取索引 1 处的元素
    print("Element at index 1:", element) // 打印 15

    try list.update(at: 0, with: 100) // 修改索引 0 处的元素为 100
    list.printList() // 打印 [100, 15, 30]

    list.clear() // 清空顺序表
    print("List after clearing:")
    list.printList() // 打印 []
} catch {
    print("Error:", error)
}




