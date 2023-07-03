//: A UIKit based Playground for presenting user interface
  
import UIKit
import PlaygroundSupport

class ListNode {
    var value: Int

    public init(_ value: Int) {
        self.value = value
    }
}

class ArrayList {
    // 存储数据元素的数组
    var values: [ListNode]?
    
    // 表的长度
    var length: Int {
        get {
            if let count = values?.count {
                return count
            }
            return 0
        }
    }
        
    // MARK: - 创建
    init() {}
    
    func insert(_ value: Int, _ index: Int) {
        let newNode = ListNode(value)
        guard let nodes = values else {
            values = []
            values?.append(newNode)
            return
        }
        
        // index 越界
        if index < 0 || index >= nodes.count {
            return
        }
        
        var temp: [ListNode] = []
        
        for i in 0 ..< nodes.count {
            if i == index {
                // 多执行一次插入
                temp.append(newNode)
            }
            temp.append(nodes[i])
        }
        values = temp
    }
    
    func remove(_ index: Int) {
        guard let nodes = values else {
            return
        }
        if index < 0 || index > (nodes.count - 1) { return }
        
        var temp: [ListNode] = []
        for i in 0 ..< nodes.count {
            if i != index {
                temp.append(nodes[i])
            }
        }
        values = temp
    }
    
    func node(at index: Int) -> ListNode? {
        guard let nodes = values else {
            return nil
        }
        if index >= 0 && index < nodes.count {
            return nodes[index]
        }
        
        return nil
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

var list = ArrayList.init()
list.insert(10, 0)
var value = list.node(at: 0)
list.remove(0)



