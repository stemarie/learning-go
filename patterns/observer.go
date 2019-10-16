package main

import (
	"fmt"
)

type IObserver interface {
	Update(state string)
}

type TextObserver struct {
	_name string
}

func (t TextObserver) Update(state string) {
	fmt.Println(t._name + ": " + state)
}

type TestSubject struct {
	_observers []IObserver
}

func (ts *TestSubject) Attach(observer IObserver) {
	ts._observers = append(ts._observers, observer)
}

func (ts *TestSubject) Detach(observer IObserver) {
	index := 0
	for i := range ts._observers {
		if ts._observers[i] == observer {
			index = i
			break
		}
	}
	ts._observers = append(ts._observers[0:index], ts._observers[index+1:]...)
}

func (ts TestSubject) Notify(state string) {
	for _, observer := range ts._observers {
		observer.Update(state)
	}
}

type TextEdit struct {
	TestSubject
	Text string
}

func (te TextEdit) SetText(text string) {
	te.Text = text
	te.TestSubject.Notify(text)
}

func main() {
	observer1 := TextObserver{"IObserver #1"}
	observer2 := TextObserver{"IObserver #2"}

	textEdit := TextEdit{}
	textEdit.Attach(observer1)
	textEdit.Attach(observer2)

	textEdit.SetText("test text")
}
