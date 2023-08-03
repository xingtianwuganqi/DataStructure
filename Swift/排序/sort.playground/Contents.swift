import UIKit

var arr = [1,9,5,2,7,8,4,3]


/*
1.比较相邻的元素。如果第一个比第二个大，就交换他们两个。
2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
3.针对所有的元素重复以上的步骤，除了最后一个。
4.持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

*/
func bubbleSort(arr: inout [Int]) {
    for i in 0..<arr.count {
        for j in 0..<(arr.count - 1 - i) {
            if arr[j] > arr[j+1] {
                arr.swapAt(j, j+1)
            }
        }
    }
//    print(arr)
}
bubbleSort(arr: &arr)


/*
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
重复第二步，直到所有元素均排序完毕。
*/
func selectSort(arr: inout [Int]) {
    for i in 0..<arr.count {
        var minIndex = i
        for j in i+1..<arr.count {
            if arr[minIndex] > arr[j] {
                minIndex = j
            }
        }
        arr.swapAt(i, minIndex)
    }
//    print(arr)
}
selectSort(arr: &arr)


// 插入排序
/*
 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
 */
func insertionSort(arr: inout [Int]) {
    for i in 1..<arr.count {
        let key = arr[i]
        var j = i - 1
        while j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j-=1
        }
        arr[j+1] = key
    }
    print(arr)
}

insertionSort(arr: &arr)




