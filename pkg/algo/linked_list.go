package algo

import (
	"dijkstra/pkg/vertices"
)

//THE FOLLOWING FILE IS BASED FROM GO AUTHORS EDITED MINORLY AND LAZILY TO SUIT MY NEEDS
//https://golang.org/src/container/list/list.go?m=text
//AVOID USING MINE AS A TEMPLATE AS I REMOVED MOST SAFETIES (that's why they are
// all private now)

// element is an element of a linked list.
type element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *element

	// The list to which this element belongs.
	list *linkedList

	// The value stored with this element.
	Value *vertices.Vertex
}

// linkedList represents a doubly linked list.
// The zero value for linkedList is an empty list ready to use.
type linkedList struct {
	root  element // sentinel list element, only &root, root.prev, and root.next are used
	len   int     // current list length excluding (this) sentinel element
	short bool
}

func linkedListNewShort() Queue {
	return Queue(new(linkedList).init(true))
}

func linkedListNewLong() Queue {
	return Queue(new(linkedList).init(false))
}

func (l *linkedList) PushOrdered(v *vertices.Vertex) {
	l.pushOrdered(v)
}

func (l *linkedList) PopOrdered() *vertices.Vertex {
	if l.short {
		return l.popBack()
	}
	return l.popFront()
}

func (l *linkedList) Len() int {
	return l.len
}

func (l *linkedList) init(short bool) *linkedList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	l.short = short
	return l
}

// front returns the first element of list l or nil.
func (l *linkedList) front() *element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// popFront pops the vertices.Vertex off the front of the list
func (l *linkedList) popFront() *vertices.Vertex {
	e := l.front()
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// popFront pops the vertices.Vertex off the front of the list
func (l *linkedList) popBack() *vertices.Vertex {
	e := l.back()
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// back returns the last element of list l or nil.
func (l *linkedList) back() *element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero linkedList value.
func (l *linkedList) lazyInit() {
	if l.root.next == nil {
		l.init(l.short)
	}
}

// pushOrdered pushes the value into the linked list in the correct position
// (ascending)
func (l *linkedList) pushOrdered(v *vertices.Vertex) *element {
	l.lazyInit()
	if l.len == 0 {
		return l.pushFront(v)
	}
	back := l.back()
	if back.Value.Distance < v.Distance {
		return l.insertValue(v, l.root.prev)
	}
	current := l.front()
	for current.Value.Distance < v.Distance && current.Value.ID != v.ID { //don't need to check if current=back cause already checked
		current = current.next
	}
	if current.Value.ID == v.ID {
		return current
	}
	return l.insertValue(v, current.prev)
}

// insert inserts e after at, increments l.len, and returns e.
func (l *linkedList) insert(e, at *element) *element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&element{Value: v}, at).
func (l *linkedList) insertValue(v *vertices.Vertex, at *element) *element {
	return l.insert(&element{Value: v}, at)
}

// remove removes e from its list, decrements l.len, and returns e.
func (l *linkedList) remove(e *element) *element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}

// pushFront inserts a new element e with value v at the front of list l and returns e.
func (l *linkedList) pushFront(v *vertices.Vertex) *element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}
