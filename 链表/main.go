package main

import "fmt"

type ListNode struct {
	val string
	next *ListNode
}
type List struct {
	Head *ListNode
}

var list *List

func init() {
	head := &ListNode{val: "0"}
	n1 := &ListNode{val: "1"}
	n2 := &ListNode{val: "2"}
	n3 := &ListNode{val: "3"}
	n4 := &ListNode{val: "4"}
	n5 := &ListNode{val: "5"}
	n6 := &ListNode{val: "6"}
	head.next = n1
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	n5.next = n6
	list = &List{
		Head: head,
	}
}

// print
func (list *List) ListPrint() string {
	str := ""
	if list.Head == nil {
		return str
	} else {
		str = list.Head.val
	}

	node := list.Head.next
	for node != nil {
		str = str + " -> " + node.val
		node = node.next
	}
	return str
}
// 插入
// 删除
// 获取
// 更新

// 两两交换
// 注释以0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6这个举例
func (list *List) swapPairs() {
	head := list.Head
	// dummy 虚拟头节点
	// 为了不对头节点单独处理，新间虚拟头节点
	dummy := &ListNode{
		next: head,
	}
	pre := dummy // pre标识头节点的前一位，便于统一处理
	for head != nil && head.next != nil { // 因为两两交换，奇数位和偶数位都需要判断
		pre.next = head.next // 先修改pre的下一个节点的指向，即pre -> 1
		next := head.next.next // 一个交换过程中的临时变量，记录交换的两个值的前一个值最终的指向，如交换0->1, 0会指向next，即2
		head.next.next = head // 修改head.next的指向，指回head，说白了，如交换0->1，则会修改0->1为0<-1，即一个翻转
		head.next = next  // 实际是修改head的指向，即指为刚刚保存的临时变量，即0->2
		pre = head  // 修改pre为下一个要调整的两个节点第一个的前一个节点
		head = next // 修改head为下一个要调整的两个节点第一个节点
	}
	list.Head = dummy.next  // 去掉虚拟头节点
}

func main() {
	fmt.Println(list.ListPrint())
	list.swapPairs()
	fmt.Println(list.ListPrint())
}
