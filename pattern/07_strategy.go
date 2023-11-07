package main

import "fmt"

// Context
type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func InitCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.Evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.storage, key)
}

func (c *Cache) Evict() {
	c.evictionAlgo.Evict(c)
	c.capacity--
}

// Strategy

type EvictionAlgo interface {
	Evict(c *Cache)
}

// Concrete algo 1
type lfu struct {
}

func NewLfu() *lfu {
	return &lfu{}
}

func (l *lfu) Evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}

// Concrete algo 2
type lru struct {
}

func NewLru() *lru {
	return &lru{}
}

func (l *lru) Evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}

// Concrete algo 3
type fifo struct {
}

func NewFifo() *fifo {
	return &fifo{}
}

func (l *fifo) Evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy")
}

func main() {
	lfu := NewLfu()
	cache1 := InitCache(lfu)
	cache1.Add("a", "1")
	cache1.Add("b", "2")
	cache1.Add("c", "3")
	lru := NewLru()
	cache1.SetEvictionAlgo(lru)
	cache1.Add("d", "4")
	fifo := NewFifo()
	cache1.SetEvictionAlgo(fifo)
	cache1.Add("e", "5")
}
