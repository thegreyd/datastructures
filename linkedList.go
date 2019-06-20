package datastructures

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

// LinkedList is a linked list
type LinkedList struct {
	head *node
	tail *node
}

var errIndexOutOfRange = errors.New("index out of range")

func (l *LinkedList) Append(data int) {
	newNode := node{data: data}

	// check if list is empty
	// if yes point head and tail to the first element
	if l.head == nil {
		l.head = &newNode
		l.tail = l.head
		return
	}

	// update previous tail to point at new node
	// and make new node current tail
	l.tail.next = &newNode
	l.tail = &newNode
}

func (l *LinkedList) Get(index int) (int, error) {
	if l.head == nil {
		return 0, errIndexOutOfRange
	}

	cursor := 0
	temp := l.head

	for cursor != index {
		if temp == nil {
			return 0, errIndexOutOfRange
		}

		temp = temp.next
		cursor++
	}

	return temp.data, nil
}

func (l *LinkedList) String() string {
	var s string

	if l.head == nil {
		return s
	}

	temp := l.head
	s = fmt.Sprintf("%v", temp.data)
	temp = temp.next

	for temp != nil {
		s = fmt.Sprintf("%v, %v", s, temp.data)
		temp = temp.next
	}
	return s
}
