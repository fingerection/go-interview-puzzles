package main

import (
	"fmt"
)

const (
	SIZE = 6
)

type Node struct {
	Key  string
	Data string
	Next *Node
}

type Cache struct {
	capacity int
	header   *Node
	d        map[string]*Node
}

func (c *Cache) Init(capacity int) {
	c.capacity = capacity
	c.d = make(map[string]*Node, 100)
	c.header = nil
}

func (c *Cache) Get(key string) *string {
	fmt.Println("get")
	if c.d[key] != nil {
		// move node to header of link
		if c.d[key] != c.header {
			p := c.header
			for ; p != nil; p = p.Next {
				if p.Next == c.d[key] {
					break
				}
			}
			p.Next = c.d[key].Next
			c.d[key].Next = c.header
			c.header = c.d[key]
		}
		return &c.d[key].Data
	}
	return nil
}

func (c *Cache) Set(key string, value string) {
	if c.d[key] != nil {
		c.d[key].Data = value
		return
	}
	if len(c.d) == c.capacity {
		// delete last pointer
		p := c.header
		for ; p.Next.Next != nil; p = p.Next {
		}
		delete(c.d, p.Next.Key)
		p.Next = nil
	}
	node := Node{Key: key, Data: value}
	c.d[key] = &node
	node.Next = c.header
	c.header = &node
}

func main() {
	fmt.Println("Start")
	c := Cache{}
	c.Init(SIZE)

	klist := []string{"k1", "k2", "k3", "k4", "k5", "k6"}
	vlist := []string{"v1", "v2", "v3", "v4", "v5", "v6"}

	for i := 0; i < len(klist); i++ {
		c.Set(klist[i], vlist[i])
	}

	p := c.Get("k1")
	if p != nil {
		fmt.Println(*p)
	}
}
