package simpledatastructures

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	newNode := newNode(5)

	if newNode.next != nil {
		t.Errorf("new created node next value should be pointer to nil")
	}

	if newNode.val != 5 {
		t.Errorf("missmatch value")
	}
}

func TestNewLinkedList(t *testing.T) {
	newLinkedList := NewLinkedList[int]()

	if newLinkedList.head != nil {
		t.Errorf("new created linkedlist head value should be pointer to nil")
	}

	if newLinkedList.size != 0 {
		t.Errorf("new created linkedlist size should be equals to 0")
	}
}

func TestSize(t *testing.T) {
	testClass := NewLinkedList[int]()

	if testClass.Size() != 0 {
		t.Errorf("new created linkedlist size should be equal 0")
	}

	testClass.Add(1)
	testClass.Add(2)

	if testClass.Size() == 0 {
		t.Errorf("after add elements to linkedlist size should not equal 0")
	}
}

func TestEmpty(t *testing.T) {
	testClass := NewLinkedList[int]()

	if !testClass.Empty() {
		t.Errorf("new created linkedlist should be empty")
	}

	testClass.Add(1)
	testClass.Add(2)

	if testClass.Empty() {
		t.Errorf("linkedlist should not be empty after added some elements")
	}
}

func TestClear(t *testing.T) {
	testClass := NewLinkedList[int]()

	testClass.Add(1)
	testClass.Add(2)
	testClass.Add(3)

	testClass.Clear()

	if !testClass.Empty() || testClass.head != nil {
		t.Errorf("linkedlist is not empty after clear operation")
	}
}

func TestAdd(t *testing.T) {
	testClass := NewLinkedList[int]()
	size := 10

	wants := []int{}
	for i := 0; i < size; i++ {
		testClass.Add(i * i)
		wants = append(wants, i*i)
	}

	if testClass.Size() != 10 {
		t.Errorf("after add 10 elements, size should equal 10, result size: %d", testClass.Size())
	}

	for i := 0; i < size; i++ {
		if testClass.Get(i) != wants[i] {
			t.Errorf("mismatch value on index %d", i)
		}
	}
}

func TestAddTo(t *testing.T) {
	testClass := NewLinkedList[int]()
	size := 10

	wants := []int{}
	for i := 0; i < size; i++ {
		testClass.AddTo(i*i, i)
		wants = append(wants, i*i)
	}

	if testClass.Size() != 10 {
		t.Errorf("after add 10 elements, size should equal 10, result size: %d", testClass.Size())
	}

	testClass.AddTo(666, 5)
	if testClass.Get(5) != 666 || testClass.Get(6) != wants[5] {
		t.Errorf("addTo does not add value at given index")
	}
}

func TestGet(t *testing.T) {
	testClass := NewLinkedList[int]()
	size := 10

	for i := 0; i < size; i++ {
		testClass.Add(i * i)
		if testClass.Get(i) != i*i {
			t.Errorf("mismatch value on index %d", i)
		}
	}
}

func TestGetFirst(t *testing.T) {
	testClass := NewLinkedList[int]()
	size := 10
	var firstVal int

	for i := 0; i < size; i++ {
		if i == 0 {
			firstVal = i * i
		}

		testClass.Add(i * i)
	}

	if testClass.GetFirst() != firstVal {
		t.Errorf("get first operation return not correct value")
	}
}

func TestGetLast(t *testing.T) {
	testClass := NewLinkedList[int]()
	size := 10
	var lastValue int

	for i := 0; i < size; i++ {
		if i == size-1 {
			lastValue = i * i
		}

		testClass.Add(i * i)
	}

	if testClass.GetLast() != lastValue {
		t.Errorf("get last return not correct value")
	}
}

func TestRemove(t *testing.T) {
	testClass := NewLinkedList[int]()
	wants := []int{}
	size := 10

	for i := 0; i < size; i++ {
		testClass.Add(i * i)
		wants = append(wants, i*i)
	}

	for i := 0; i < size/2; i++ {
		testClass.Remove(0)
	}

	if testClass.GetLast() != wants[size-1] ||
		testClass.GetFirst() != wants[size/2] {
		t.Errorf("after remove operation list contains not valid data")
	}

	if testClass.Size() != size/2 {
		t.Errorf("after remove operation list size is not correct")
	}
}

func TestPop(t *testing.T) {
	testClass := NewLinkedList[int]()
	wants := []int{}
	size := 10

	for i := 0; i < size; i++ {
		testClass.Add(i * i)
		wants = append(wants, i*i)
	}

	lastVal := testClass.GetLast()
	firstVal := testClass.GetFirst()

	if testClass.Pop(0) != firstVal || testClass.Pop(testClass.Size()-1) != lastVal {
		t.Errorf("pop operations return not valid data")
	}

	if testClass.Size() != size-2 {
		t.Errorf("after pop operations size of structure is not correct")
	}
}

func TestGetNode(t *testing.T) {
	testClass := NewLinkedList[int]()

	testClass.Add(5)
	testClass.Add(5)
	testClass.Add(5)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("node should panic, if index value is less than 0 or greather than size of collection")
		}
	}()

	testClass.getNode(-1)
	testClass.getNode(3)
}

func TestGetLastNode(t *testing.T) {
	testClass := NewLinkedList[int]()

	testClass.Add(1)
	testClass.Add(2)
	testClass.Add(3)

	if testClass.getLastNode().val != 3 {
		t.Errorf("missmatch value")
	}

	if testClass.getLastNode().next != nil {
		t.Errorf("getLastNode does not return last node")
	}
}
