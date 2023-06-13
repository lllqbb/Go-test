package main

import (
	"fmt"
)

// 节点
type Node struct {
	Value int   // 当前值
	Next  *Node // 下个节点的地址
}

// 初始化头结点
var head = new(Node)

// 添加节点
func addNode(t *Node, v int) int {
	if head == nil {
		t = &Node{v, nil}
		head = t
		return 0
	}

	if t.Value == v {
		fmt.Println("当前节点已存在:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}

// 遍历链表
func traverse(t *Node) {
	if t == nil {
		fmt.Println("->空链表！")
	}

	for t != nil {
		fmt.Println("%d->", t.Value)
		t = t.Next
	}

	fmt.Println("---------遍历结束--------")
}

func lookupNode(t *Node, v int) bool {
	if head == nil {
		t = &Node{v, nil}
		head = t
		return false
	}

	if t.Value == v {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func main() {
	fmt.Println("测试")
}
