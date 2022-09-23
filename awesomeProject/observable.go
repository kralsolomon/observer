package main

import "fmt"

type Observable struct {
	name        string
	observers   []Observer
	temperature float32
	pressure    float32
}

func (Observable *Observable) Subscribe(observer Observer) {
	Observable.observers = append(Observable.observers, observer)
	fmt.Printf("%v has been added %v \n", Observable.name, observer.name)
}

func (Observable *Observable) Unsubscribe(observer Observer) {
	for i := range Observable.observers {
		if Observable.observers[i] == observer {
			Observable.observers = append(Observable.observers[:i], Observable.observers[i+1:]...)
		}
	}
}

func (Observable *Observable) NotifyAll() {
	for i := range Observable.observers {
		Observable.observers[i].Update(Observable.observers[i])
	}
}

func (Observable *Observable) SetTemperature(newTemp float32) {
	Observable.temperature = newTemp
	Observable.NotifyAll()
}

func (Observable *Observable) SetPressure(newPressure float32) {
	Observable.pressure = newPressure
	Observable.NotifyAll()
}

func NewSubject(name string, initialTemperature float32, initialPressure float32) *Observable {
	var object = new(Observable)
	object.name = name
	object.pressure = initialTemperature
	object.pressure = initialPressure
	object.observers = []Observer{}
	return object
}
