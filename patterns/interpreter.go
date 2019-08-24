package main

import (
	"fmt"
)

// IExpression is AbstractExpression
type IExpression interface {
	Interpret(i int) bool
}

// DevExpression is TerminalExpression
type DivExpression struct {
	divider int
}

// Interpret DivExpression method
func (e DivExpression) Interpret(i int) bool {
	return i%e.divider == 0
}

// OrExpression is NonterminalExpression
type OrExpression struct {
	exp1 IExpression
	exp2 IExpression
}

// Interpret OrExpression method
func (e OrExpression) Interpret(i int) bool {
	return e.exp1.Interpret(i) || e.exp2.Interpret(i)
}

// AndExpression is NonterminalExpression
type AndExpression struct {
	exp1 IExpression
	exp2 IExpression
}

// Interpret AndExpression method
func (e AndExpression) Interpret(i int) bool {
	return e.exp1.Interpret(i) && e.exp2.Interpret(i)
}

func main() {
	divExp5 := DivExpression{5}
	divExp7 := DivExpression{7}
	orExp := OrExpression{divExp5, divExp7}
	andExp := AndExpression{divExp5, divExp7}

	result1 := orExp.Interpret(21)
	result2 := andExp.Interpret(21)
	result3 := andExp.Interpret(35)

	fmt.Println("21 can be divided by 7 or 5", result1)
	fmt.Println("21 can be divided by 7 and 5", result2)
	fmt.Println("35 can be divided by 7 and 5", result3)
}
