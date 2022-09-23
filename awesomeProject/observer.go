package main

import "fmt"

type Observer struct {
	name string
}

func NewListener(name string) *Observer {
	var listener = new(Observer)
	listener.name = name
	return listener
}

func (Observer *Observer) Update(observer Observer) {
	fmt.Printf("%v detected that the weather has changed \n", observer.name)
}
