package main

import "fmt"

func main() {
	deque := NewDeque(100)
	deque.PushBack(10)
	deque.PushBack(15)
	deque.PushBack(20)
	deque.PushBack(1)

	fmt.Println(deque.PopBack())
	fmt.Println(deque.PopFront())

	fmt.Println(deque.Back())
	fmt.Println(deque.Front())

	fmt.Println(deque.PopBack())
	fmt.Println(deque.PopBack())

	fmt.Println(deque.Front())
}

type Value int

type Deque interface {
	PushBack(val Value)
	PushFront(val Value)

	PopBack() Value
	PopFront() Value

	Front() Value
	Back() Value
}

func NewDeque(capacity uint) Deque {
	return &deque{
		circle:   make([]Value, capacity),
		capacity: capacity,
		size:     0,
		front:    capacity - 1,
		back:     0,
	}
}

// front ><=========>< back
// 0 1 2 3 4 5 6
type deque struct {
	circle   []Value
	capacity uint
	size     uint
	front    uint
	back     uint
}

func (d *deque) decrementIndex(i *uint) {
	*i = (d.capacity + *i - 1) % d.capacity
}

func (d *deque) incrementIndex(i *uint) {
	*i = (*i + 1) % d.capacity
}

func (d *deque) incremented(i uint) uint {
	return (i + 1) % d.capacity
}

func (d *deque) decremented(i uint) uint {
	return (d.capacity + i - 1) % d.capacity
}

func (d *deque) PushBack(val Value) {
	if d.size == d.capacity {
		panic("Out of memory!")
	}

	d.size++
	d.circle[d.back] = val
	d.incrementIndex(&d.back)
}

func (d *deque) PushFront(val Value) {
	if d.size == d.capacity {
		panic("Out of memory!")
	}

	d.size++
	d.circle[d.front] = val
	d.decrementIndex(&d.front)
}

func (d *deque) PopBack() Value {
	if d.size == 0 {
		panic("Deque is already empty!")
	}

	d.size--
	d.decrementIndex(&d.back)
	return d.circle[d.back]
}

func (d *deque) PopFront() Value {
	if d.size == 0 {
		panic("Deque is already empty!")
	}

	d.size--
	d.incrementIndex(&d.front)
	return d.circle[d.front]
}

func (d *deque) Back() Value {
	if d.size == 0 {
		panic("Deque is empty!")
	}

	return d.circle[d.decremented(d.back)]
}

func (d *deque) Front() Value {
	if d.size == 0 {
		panic("Deque is empty!")
	}

	return d.circle[d.incremented(d.front)]
}
