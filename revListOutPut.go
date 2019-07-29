package main

import "fmt"

type node struct {
	val  int
	next *node
}

//链表翻转数据
func main() {
	n3 := node{val: 4, next: nil}
	n2 := node{val: 3, next: &n3}
	fmt.Println(n2.val)

	n1 := node{val: 2, next: &n2}
	head := node{val: 1, next: &n1}
	revOutPut(&head)
}

func revOutPut(head *node) {
	//解释一下这里为什么要把地址返回回来，否则head还是不变
	//首先go参数都是传值的
	//这里的head的值是一个地址
	//在reverseList函数内的head是一个局部变量，值也是head的地址，
	//在reverseList函数内改变head的地址并不会影响外面head的地址
	//但是如果在reverseList函数内直接改变head指向的内容，则真正的内容会被改变
	head = reverseList(head)
	tmp := head
	for tmp != nil {
		fmt.Println(tmp.val)
		tmp = tmp.next
	}
}

/**
翻转链表
*/
func reverseList(head *node) *node {
	if head == nil {
		return nil
	}
	var prev *node
	cur := head
	next := head.next
	for next != nil {
		cur.next = prev
		if next.next == nil {
			next.next = cur
			head = next
			break
		}
		prev = cur
		cur = next
		next = next.next
	}
	return head
}
