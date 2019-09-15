package main

import (
	"fmt"
)

type IStrategy interface {
	DoOperation(a int, b int) int
}

//Concrete Strategy
type AddStrategy struct{}

func (s AddStrategy) DoOperation(a int, b int) int {
	return a + b
}

//Concrete Strategy
type SubstractStrategy struct{}

func (s SubstractStrategy) DoOperation(a int, b int) int {
	return a - b
}

type Calc struct {
	_strategy IStrategy
}

func (c Calc) Execute(a int, b int) int {
	if c._strategy == nil {
		return 0
	}
	return c._strategy.DoOperation(a, b)
}

func (c *Calc) SetStrategy(strategy IStrategy) {
	c._strategy = strategy
}

func main() {
	// Client
	calc1 := Calc{}
	result1 := calc1.Execute(5, 3)
	fmt.Println(result1)

	calc1.SetStrategy(AddStrategy{})
	result2 := calc1.Execute(5, 3)
	fmt.Println(result2)

	calc1.SetStrategy(SubstractStrategy{})
	result3 := calc1.Execute(5, 3)
	fmt.Println(result3)
}
