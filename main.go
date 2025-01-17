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
	Queue     Queue // Doubly linked list
	Hash      Hash  // Map for lookup for items
	Hits      int
	Miss      int
	Evictions int
}

// stores the cache items
type Hash map[string]*Node

// The hash map provides O(1) lookup time for cache items.
// Key is value of the item and value is the pointer to the corresponding node

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
		c.Hits++
	} else {
		node = &Node{Value: str}
		c.Miss++
	}

	//If the item already exists in the cache(Hash),
	// It is moved to the front of the queue.
	// If the item does not exist, it is added to the front of the queue.
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	// The item is removed from the queue by updating the pointers of its neighboring nodes.
	fmt.Printf("The value removed is %s\n", n.Value)

	//getting the node neighbors
	left := n.Left
	right := n.Right

	//Bypassed the node to be removed
	left.Right = right
	right.Left = left

	c.Queue.Length -= 1

	delete(c.Hash, n.Value) // The item is also removed from the Hash map.
	c.Evictions++

	return n
}

func (c *Cache) Add(n *Node) *Node {
	fmt.Printf("The value added is %s\n", n.Value)

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n // Point head's next to the new node
	n.Left = c.Queue.Head  // Point new node's left to head
	n.Right = tmp          // Point new node's right to current first node
	tmp.Left = n           // Update current first node's left to new node

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

func (c *Cache) Clear() {
	c.Queue = NewQueue()
	c.Hash = Hash{}
	fmt.Println("Cache cleared.")
}

// Stats displays usage statistics for the cache
func (c *Cache) Stats() {
	fmt.Printf("Cache Stats - Hits: %d, Misses: %d, Evictions: %d\n", c.Hits, c.Miss, c.Evictions)
}
func main() {
	fmt.Println("Starting cache")

	// Cache is initialized with empty queue and empty hashmap
	cache := NewCache()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Commands:")
	fmt.Println("  add <value>   - Add a value to the cache")
	fmt.Println("  stats         - Show cache statistics")
	fmt.Println("  clear         - Clear the cache")
	fmt.Println("  exit          - Exit the program")

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if strings.ToLower(input) == "exit" {
			break
		}
		if strings.ToLower(input) == "stats" {
			cache.Stats()
		}
		if strings.ToLower(input) == "clear" {
			cache.Clear()
		}
		// Process the input
		cache.Check(input) // Check if the value already exists in the queue or not
		cache.Display()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
