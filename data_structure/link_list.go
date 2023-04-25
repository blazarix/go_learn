package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

// 获取指定位置的节点
func (n *Node) get(index int) *Node {
	p := n
	for i := 0; i < index; i++ {
		p = p.Next
		if p == nil {
			break
		}
	}
	return p
}

// 获取尾部节点
func (n *Node) getTail() *Node {
	// 遍历获取尾部节点
	p := n
	for {
		if p.Next == nil {
			break
		}
		p = p.Next
	}
	return p
}

// 添加一个值在链表头部,新的值成为头部
func (n *Node) addHead(val interface{}) {
	// 拷贝原来的n
	oldHead := new(Node)
	oldHead.Value = n.Value
	oldHead.Next = n.Next

	// 修改头节点的值，指向拷贝的原头节点
	n.Value = val
	n.Next = oldHead
	//head := new(Node)
	//head.Value = val
	//head.Next = n
}

// 添加一个值在链表末尾，新的值成为尾部
func (n *Node) addTail(value interface{}) {
	// 获取原尾部节点
	preTail := n.getTail()

	tail := new(Node)
	tail.Value = value
	preTail.Next = tail
}

// 添加一个值在指定位置
func (n *Node) addIndex(index int, value interface{}) {
	// 找到原位置的上一个节点
	node := n.get(index - 1)
	newNode := new(Node)
	newNode.Value = value
	newNode.Next = node.Next
	node.Next = newNode
}

// 如果索引有效，删除指定位置节点
func (n *Node) deleteIndex(index int) int {
	// 找到原位置节点的上一个节点
	node := n.get(index - 1)
	if node == nil {
		return -1
	}
	node.Next = node.Next.Next
	return 1
}

// 遍历链表
func (n *Node) print() {
	p := n
	for {
		if p.Next == nil {
			fmt.Printf("%v", p.Value)
			break
		} else {
			fmt.Printf("%v->", p.Value)
		}
		p = p.Next
	}
}

// 如果一个字符串使用链表存储，怎么判断这个字符串是不是回文字符串（从前往后，从后往前读都一样，比如level)
//func isPalindrome(s string) bool {
//
//}

func main() {
	list := new(Node)
	list.Value = 5

	list.addHead(3)
	list.addHead(2)
	list.addHead(99)
	list.addTail(100)
	list.addIndex(2, "xi")
	fmt.Println(list.deleteIndex(100))
	fmt.Println(list.get(2).Value)

	list.print()
}
