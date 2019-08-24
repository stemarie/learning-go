package main

import (
	"fmt"
)

// ICommand interface
type ICommand interface {
	Execute()
}

// BankClient Invoker
type BankClient struct {
	putCommand ICommand
	getCommand ICommand
}

// PutMoney method for BankClient
func (bc BankClient) PutMoney() {
	bc.putCommand.Execute()
}

// GetMoney method for BankClient
func (bc BankClient) GetMoney() {
	bc.getCommand.Execute()
}

// Bank Receiver
type Bank struct{}

// GiveMoney Bank Method
func (b Bank) GiveMoney() {
	fmt.Println("money to the client")
}

// ReceiveMoney Bank Method
func (b Bank) ReceiveMoney() {
	fmt.Println("money from the client")
}

// PutCommand is the ConcreteCommand
type PutCommand struct {
	bank Bank
}

// Execute PutCommand Method
func (pc PutCommand) Execute() {
	pc.bank.ReceiveMoney()
}

// GetCommand is the ConcreteCommand
type GetCommand struct {
	bank Bank
}

// Execute GetCommand Method
func (pc GetCommand) Execute() {
	pc.bank.GiveMoney()
}

// Client
func main() {
	bank := Bank{}
	cPut := PutCommand{bank}
	cGet := GetCommand{bank}
	client := BankClient{cPut, cGet}
	client.GetMoney()
	// printed: money to the client
	client.PutMoney()
	// printed: money from the client
}
