package simpledatastructures

// node for linked list, contains value
// and next element in that structure.
type node[T any] struct {
	val  T
	next *node[T]
}

// newNode constructor initialize node with inserted
// value and pointing to nothing.
func newNode[T any](val T) *node[T] {
	return &node[T]{
		val:  val,
		next: nil,
	}
}

// linkedList contains head of structure and number of
// elements inside.
type linkedList[T any] struct {
	head *node[T]
	size int
}

// NewLinkedList constructor initialize empty linked list.
func NewLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{
		head: nil,
		size: 0,
	}
}

// Size return how many element linked list currently contains.
func (ll *linkedList[T]) Size() int {
	return ll.size
}

// Empty return if linekd list is currently empty.
func (ll *linkedList[T]) Empty() bool {
	return ll.Size() == 0
}

// Add value as a last element in linked list.
func (ll *linkedList[T]) Add(val T) {
	if ll.Size() == 0 {
		ll.head = newNode(val)
		ll.size += 1
		return
	}

	ll.getLastNode().next = newNode(val)
	ll.size += 1
}

// AddTo add value on given index.
func (ll *linkedList[T]) AddTo(val T, idx int) {
	if idx == 0 {
		curr := ll.head
		ll.head = newNode(val)
		ll.head.next = curr

		return
	}

	prev := ll.getNode(idx - 1)
	new := newNode(val)
	curr := prev.next

	prev.next = new
	new.next = curr

	ll.size += 1
}

// Get return element on given index in linked list,
// or panic if given index is out of range.
func (ll *linkedList[T]) Get(idx int) T {
	return ll.getNode(idx).val
}

// getNode return pointer to node on given index,
// or panic if given index is out of range.
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

// getLastNode return node on last index.
func (ll *linkedList[T]) getLastNode() *node[T] {
	return ll.getNode(ll.Size() - 1)
}
