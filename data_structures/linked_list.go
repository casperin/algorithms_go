package datastructures

import "fmt"

// Linked list implementation where the value in the list is int (which can be
// easily mapped to any other data type using a map[int]MyStruct).
type LinkedList struct {
	head *linkedListNode
}

type linkedListNode struct {
	value int
	next  *linkedListNode
}

// Implementation of the stringer interface. Returns something like
// 1->2->3->nil
func (l *LinkedList) String() string {
	s := ""
	node := l.head
	for node != nil {
		s += fmt.Sprintf("%v->", node.value)
		node = node.next
	}
	return s + "nil"
}

// Push a value to the front of the list.
//
// This: 1->2->3
// To:   7->1->2->3
// with Push(7)
func (l *LinkedList) Push(value int) {
	node := linkedListNode{value, l.head}
	l.head = &node
}

// Iterates over the list (front to back) and pushes the values, one by one,
// onto the list.
//
// This: 1->2
// To:   20->10->1->2
// with PushAll([]int{10, 20})
func (l *LinkedList) PushAll(values []int) {
	for _, value := range values {
		node := linkedListNode{value, l.head}
		l.head = &node
	}
}

// Pops the first values off the top of the list. Returns a second argument,
// ok, which will be false when no value was found.
func (l *LinkedList) Pop() (int, bool) {
	if l.head == nil {
		return 0, false
	}
	value := l.head.value
	l.head = l.head.next
	return value, true
}

// Like Pop(), except the value stays in the list.
func (l *LinkedList) Peek() (int, bool) {
	if l.head == nil {
		return 0, false
	}
	return l.head.value, true
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) Size() int {
	size := 0
	node := l.head
	for node != nil {
		size++
		node = node.next
	}
	return size
}

// Deletes a value from the list. Returns true if the value was found (and
// deleted). Otherwise false.
func (l *LinkedList) Delete(x int) bool {
	// First we need to check the reference to the head.
	if l.head == nil {
		return false
	}
	if l.head.value == x {
		l.head = l.head.next
		return true
	}

	node := l.head

	for node.next != nil {
		if node.next.value == x {
			// We just update the current node's ref to skip the one we just
			// found (that is, the one we want to delete).
			node.next = node.next.next
			return true
		}

		node = node.next
	}
	return false
}

// Reverses the linked list.
// This: 1->2->3
// to:	 3->2->1
func (l *LinkedList) Reverse() {
	current := l.head
	var prev *linkedListNode = nil

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev
}

// Rotates the linked list around the k'th node. The k+1'th node becomes the
// first node, and whatever was before gets attached to the end.
//
// This: 1->2->3->4->5->6
// To:   5->6->1->2->3->4
// with list.Rotate(4)
//
// If k >= number of items in the list, then Rotate is a noop.
func (l *LinkedList) Rotate(k int) {
	if k < 1 {
		return
	}

	current := l.head

	// Find kth element (or reach end of list)
	for i := 1; i < k && current != nil; i++ {
		current = current.next
	}

	// Noop if k < len(list)
	if current == nil {
		return
	}

	kthNode := current

	// Walk to the end
	for current.next != nil {
		current = current.next
	}

	current.next = l.head // attach front to the end (6 -> 1)
	l.head = kthNode.next // attach kth node to the front (head = 5)
	kthNode.next = nil    // kth node is now the end (4 -> nil)
}
