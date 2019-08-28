package main

import (
	"fmt"
)

// IIntIterator Iterator for Int
type IIntIterator interface {
	First()
	Next()
	IsDone() bool
	CurrentItem() int
}

// IIntAggregate Aggregate for Int
type IIntAggregate interface {
	GetIterator() IIntIterator
}

// Numbers bogus type for Iterator example
type Numbers struct {
	Data []int
}

// GetIterator Numbers method
func (n Numbers) GetIterator() IIntIterator {
	return &Iterator{n, 0}
}

// Iterator for Numbers type
type Iterator struct {
	_numbers Numbers
	_index   int
}

// First Iterator method
func (i *Iterator) First() {
	i._index = 0
}

// Next Iterator method
func (i *Iterator) Next() {
	i._index++
}

// IsDone Iterator method
func (i *Iterator) IsDone() bool {
	return i._index >= len(i._numbers.Data)
}

// CurrentItem Iterator method
func (i *Iterator) CurrentItem() int {
	return i._numbers.Data[i._index]
}

// main for running example
func main() {
	numbers := Numbers{[]int{2, 3, 5, 7, 11}}
	iterator := numbers.GetIterator()
	sum := 0
	for iterator.First(); !iterator.IsDone(); iterator.Next() {
		sum += iterator.CurrentItem()
	}
	fmt.Println("Sum is", sum)
}
