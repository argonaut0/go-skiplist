package main

import (
	"fmt"
	"math/rand"
)

type SkipListNode struct {
	Prev []*SkipListNode
	Next []*SkipListNode
	Key  int
}

func NewSkipList() *SkipListNode {
	return &SkipListNode{
		Prev: make([]*SkipListNode, 0),
		Next: make([]*SkipListNode, 0),
		Key:  -1,
	}
}

func (s *SkipListNode) Print() {
	curr := s
	if s.Key == -1 {
		if len(s.Next) == 0 {
			return
		}
		curr = curr.Next[0]
	}
	fmt.Println("List:")
	for curr != nil {
		for range len(curr.Prev) {
			fmt.Printf("%v\t", curr.Key)
		}
		if len(curr.Next) != len(curr.Prev) {
			panic(fmt.Sprintf("Length of next and prev are not equal %v", curr))
		}
		fmt.Print("\n")
		curr = curr.Next[0]
	}
}

func (s *SkipListNode) Search(key int) *SkipListNode {
	k := len(s.Next) - 1
	currNode := s
	for k >= 0 {
		for currNode.Next[k] != nil {
			if currNode.Next[k].Key > key {
				break
			}
			if currNode.Next[k].Key == key {
				return currNode.Next[k]
			}
			currNode = currNode.Next[k]
		}
		k = k - 1
	}
	return nil
}

// Inserts a Node and returns a pointer to that node.
func (s *SkipListNode) Insert(key int) {
	lv := level()
	newNode := SkipListNode{
		Prev: make([]*SkipListNode, lv+1),
		Next: make([]*SkipListNode, lv+1),
		Key:  key,
	}
	for len(s.Next)-1 < lv {
		s.Next = append(s.Next, nil)
	}
	k := lv
	currNode := s
	for k >= 0 {
		for currNode.Next[k] != nil {
			if currNode.Next[k].Key > key {
				currNode.Next[k].Prev[k] = &newNode
				break
			}
			currNode = currNode.Next[k]
		}
		newNode.Prev[k] = currNode
		newNode.Next[k] = currNode.Next[k]
		currNode.Next[k] = &newNode
		k = k - 1
	}
}

// Deletes the node with a given key. Returns true if an element was deleted
func (s *SkipListNode) Delete(key int) bool {
	e := s.Search(key)
	if e != nil {
		for i := range e.Prev {
			if e.Prev[i] != nil {
				e.Prev[i].Next[i] = e.Next[i]
			}
			if e.Next[i] != nil {
				e.Next[i].Prev[i] = e.Prev[i]
			}
		}
		return true
	}
	return false
}

// Returns an integer X~Geometric(1/2)
func level() int {
	l := 0
	for rand.Intn(2) > 0 {
		l++
	}
	return l
}
