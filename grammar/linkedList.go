package main

type MyLinkedList struct {
	val  int
	Next *MyLinkedList
}

// Constructor /** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		val: 0, Next: nil,
	}
}

// Get /** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {

	cur := this
	count := 0

	for cur != nil {
		if count == index {
			return cur.val
		}
		cur = cur.Next
		count += 1
	}

	return -1
}

// AddAtHead /** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {

	if this.val == 0 && this.Next == nil {
		this.val = val
		return
	}

	curVal := this.val
	this.val = val
	this.Next = &MyLinkedList{
		curVal, this.Next,
	}
}

// AddAtTail /** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {

	if this.val == 0 && this.Next == nil {
		this.val = val
		return
	}
	cur := this
	if this.Next != nil {
		cur = cur.Next
	}

	cur.Next = &MyLinkedList{val, nil}
}

// AddAtIndex /** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	count := 0
	cur := this
	for cur.Next != nil {
		if count != index-1 {
			cur.Next = &MyLinkedList{val, cur.Next}
			return
		}

		cur = cur.Next
		count += 1
	}

	this.AddAtTail(val)
}

// DeleteAtIndex /** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.Next != nil {
			this.val = this.Next.val
			this.Next = this.Next.Next
			return
		} else {
			this = &MyLinkedList{}
		}
	}
	count := 0
	cur := this
	for cur.Next != nil {
		if index == count+1 {
			cur.Next = cur.Next.Next
			return
		}
		cur = cur.Next
		count += 1
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
