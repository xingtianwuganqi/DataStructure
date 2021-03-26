//: [Previous](@previous)

import Foundation
import PlaygroundSupport

class Node {
    var value: Int = 0
    var next: Node?
    init(_ value: Int, _ next: Node?) {
        self.value = value
        self.next = next
    }
}

class LinkList {
    var head : Node?
    var tail : Node?
    
    func append(node: Node) {
        if let tailNode = tail {
            tailNode.next = node
        }else{
            head = node
        }
        tail = node
    }
    
    func nodeAt(_ index: Int) -> Node? {
        if index >= 0 {
            var node = head
            var i = index
            while node != nil {
                if i == 0 {
                    return node
                }
                i-=1
                node = node!.next
            }
        }
        return nil
    }
}

var list = LinkList.init()
list.append(node: Node.init(1,nil))
var value = list.nodeAt(0)
print(value)


