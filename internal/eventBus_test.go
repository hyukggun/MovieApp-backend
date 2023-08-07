package internal

import (
	"fmt"
	"testing"
)

func TestCreateEventBus(t *testing.T) {
	bus := CreateEventBus()
	defer bus.Dispose()
	if bus == nil {
		t.Error("event bus is nil")
	}
}

func TestAddSubscription(t *testing.T) {
	bus := CreateEventBus()
	defer bus.Dispose()
	bus.AddSubscription(func(n int) { fmt.Println(n * n) })
	if len(bus.Subscriptions) != 1 {
		t.Error("add subscription failed")
	}
}

func TestNext(t *testing.T) {
	bus := CreateEventBus()
	defer bus.Dispose()
	bus.AddSubscription(func(n int) {
		if n%2 == 1 {
			t.Errorf("%d is not divided by 2", n)
		} else {
			t.Logf("%d is divided by 2", n)
		}
	})
	bus.Next(2)
	bus.Next(4)
	bus.Next(6)
	bus.Next(7)
	bus.Next(8)
	bus.Next(9)
	bus.Next(10)
	bus.Next(11)
}
