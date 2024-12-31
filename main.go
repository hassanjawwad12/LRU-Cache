package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const SIZE = 5

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

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("The value removed is %s\n", n.Value)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length -= 1

	delete(c.Hash, n.Value)
	return n
}

func (c *Cache) Add(n *Node) *Node {
	fmt.Printf("The value added is %s\n", n.Value)

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n //new item is added at head
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++

	//if size is full remove the last element to accomodate the incoming one
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
