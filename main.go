package main

import (
	"fmt"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string) {
	node := &Node{}
}

func main() {
	fmt.Println("Starting cache")
	cache := NewCache()
	for _, word := range []string{"parrot", "tree", "lion", "forest", "rain"} {
		cache.Check(word) //check if the value already exists in the queue or not
		cache.Display()
	}
}
