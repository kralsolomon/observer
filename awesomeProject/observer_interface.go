package main

type observer interface {
	Subscribe(Observer Observer)
	Unsubscribe(Observer Observer)
	NotifyAll()
}
