package simpledatastructures

// node is a structure for contains value and
// next node, to create linked list.
type node[T any] struct {
	val  T
	next *node[T]
}

// newNode is constructor that initialize node with
// inserted value. newNode return node that point to
// nothing.
func newNode[T any](val T) *node[T] {
	return &node[T]{
		val:  val,
		next: nil,
	}
}

// linkedList is a structure that contains node head and
// follow size of structure.
type linkedList[T any] struct {
	head *node[T]
	size int
}

// NewLinkedList is constructor that initialize empty
// linked list.
func NewLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{
		head: nil,
		size: 0,
	}
}

// Size return how many elements list acctually contains.
func (ll *linkedList[T]) Size() int {
	return ll.size
}

// Empty return boolean value, true if linkedlist is
// empty and false if it is not.
func (ll *linkedList[T]) Empty() bool {
	return ll.Size() == 0
}

// Clear linked list, that operation remove all elements
// by setting head of the list to pointing to nothing, also
// change size of structure to zero.
func (ll *linkedList[T]) Clear() {
	ll.head = nil
	ll.size = 0
}

// Add value as a last element in structure, after that
// operation should increase size by one.
func (ll *linkedList[T]) Add(val T) {
	if ll.Empty() {
		ll.head = newNode(val)
		ll.size += 1
		return
	}

	ll.getLastNode().next = newNode(val)
	ll.size += 1
}

// AddTo add value to given index inside structure,
// panic if given index is out of range, that operation
// should increase size by one.
func (ll *linkedList[T]) AddTo(val T, idx int) {
	if idx == 0 {
		curr := ll.head
		ll.head = newNode(val)
		ll.head.next = curr

		ll.size += 1

		return
	}

	prev := ll.getNode(idx - 1)
	new := newNode(val)
	curr := prev.next

	prev.next = new
	new.next = curr

	ll.size += 1
}

// Get value from node at given index. That operation
// does not change size of whole structure.
func (ll *linkedList[T]) Get(idx int) T {
	return ll.getNode(idx).val
}

// GetFirst value from node at index zero. That operation
// does not change size of whole structure.
func (ll *linkedList[T]) GetFirst() T {
	return ll.Get(0)
}

// GetLast value from node at last index. That operation
// does not change size of whole structure.
func (ll *linkedList[T]) GetLast() T {
	return ll.Get(ll.Size() - 1)
}

// Remove value with node at given index. That operation
// should decrease size by one, function also panic if
// structure is empty or given index is out of range.
func (ll *linkedList[T]) Remove(idx int) {
	if ll.Empty() {
		panic("index out of range")
	}

	if idx == 0 {
		ll.head = ll.head.next
		ll.size -= 1

		return
	}

	prev := ll.getNode(idx - 1)
	curr := prev.next
	next := curr.next

	prev.next = next
	ll.size -= 1
}

// Remove value with node at given index and returns it. That
// operation should decrease size by one, function also panic if
// structure is empty or given index is out of range.
func (ll *linkedList[T]) Pop(idx int) T {
	val := ll.Get(idx)
	ll.Remove(idx)

	return val
}

// getNode return node at given index, if collection is empty
// or index is out of range, function call panic.
func (ll *linkedList[T]) getNode(idx int) *node[T] {
	if ll.Empty() || idx >= ll.Size() || idx < 0 {
		panic("index out of range")
	}

	n := ll.head
	i := 0
	for i < idx {
		n = n.next
		i += 1
	}

	return n
}

// getLastNode return node at last index, if collection is
// empty, function call panic.
func (ll *linkedList[T]) getLastNode() *node[T] {
	return ll.getNode(ll.Size() - 1)
}
