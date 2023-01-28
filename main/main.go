package main

import (
	"fmt"

	sds "github.com/pniewiarowski/simple-data-structures"
)

func main() {
	ll := sds.NewLinkedList[int]()

	fmt.Println(ll.Empty())

	ll.Add(5)
	ll.Add(6)
	ll.Add(7)
	ll.Add(8)
	ll.Add(9)

	ll.AddTo(123, 0)
	ll.AddTo(123, 3)
	ll.AddTo(123, ll.Size())

	for i := 0; i < ll.Size(); i++ {
		fmt.Println(ll.Get(i))
	}
}
