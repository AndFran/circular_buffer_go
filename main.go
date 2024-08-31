package main

import (
	"errors"
	"fmt"
)

var ErrCircularBufferEmpty = errors.New("circular buffer empty")

type CircularBuffer struct {
	Size   int
	buffer []int
	Low    int
	High   int
	Count  int
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		Size:   size,
		buffer: make([]int, size),
	}
}

func (c *CircularBuffer) isEmpty() bool {
	return c.Count == 0
}

func (c *CircularBuffer) isFull() bool {
	return c.Size == c.Count
}

func (c *CircularBuffer) Add(value int) {
	if c.isFull() {
		c.Low = (c.Low + 1) % c.Size
	} else {
		c.Count++
	}
	c.buffer[c.High] = value
	c.High = (c.High + 1) % c.Size
}

func (c *CircularBuffer) Remove() (int, error) {
	if c.isEmpty() {
		return 0, ErrCircularBufferEmpty
	}
	c.Count--
	value := c.buffer[c.Low]
	c.Low = (c.Low + 1) % c.Size
	return value, nil
}

func (c *CircularBuffer) String() string {
	return fmt.Sprintf("%+v, low: %v, high: %v", c.buffer, c.Low, c.High)
}

func main() {
	c := NewCircularBuffer(3)
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Add(4)
	fmt.Println(c.Remove())
	fmt.Println(c.Remove())
	fmt.Println(c.Remove())
	fmt.Println(c.Remove())
	fmt.Println(c)
}
