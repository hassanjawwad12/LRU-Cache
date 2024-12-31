package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const SIZE = 5

// Node and Queue initialization
// Represents an item in the cache
type Node struct {
	Value string // data stored in the node
	Left  *Node  // pointer to prev node in queue
	Right *Node  // pointer to next node in queue
}

// Represents doubly linked list
type Queue struct {
	Head   *Node // dummy node at queue beginning
	Tail   *Node // dummy node at queue end
	Length int
}

// Cache initialization
// Combines queue and hash to implement LRU cache
type Cache struct {
	Queue Queue // Doubly linked list
	Hash  Hash  // Map for lookup for items
}

// stores the cache items
type Hash map[string]*Node

// The hash map provides O(1) lookup time for cache items.
// key is value of the item and value is the pointer to the corresponding node

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

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: str}
	}

	//If the item already exists in the cache(Hash),
	// it is moved to the front of the queue.
	// If the item does not exist, it is added to the front of the queue.
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	// The item is removed from the queue by updating the pointers of its neighboring nodes.
	fmt.Printf("The value removed is %s\n", n.Value)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length -= 1

	delete(c.Hash, n.Value) // The item is also removed from the Hash map.

	return n
}

func (c *Cache) Add(n *Node) *Node {
	fmt.Printf("The value added is %s\n", n.Value)

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n // new item is added at head
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++

	// If the queue exceeds its capacity (SIZE), the least recently used item (at the tail of the queue) is removed.
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}

	return n
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Value)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Starting cache")

	// cache is initialized with empty queue and empty hashmap
	cache := NewCache()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter strings (type 'exit' to stop):")

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if strings.ToLower(input) == "exit" {
			break
		}
		// Process the input
		cache.Check(input) // Check if the value already exists in the queue or not
		cache.Display()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
