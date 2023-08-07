package internal

import (
	"sync"
)

type IntEventBus struct {
	Observable    chan int
	Subscriptions []func(int)
	wg            sync.WaitGroup
}

func CreateEventBus() *IntEventBus {
	return &IntEventBus{Observable: make(chan int, 2), Subscriptions: []func(int){}, wg: sync.WaitGroup{}}
}

func (eb *IntEventBus) AddSubscription(subscript func(int)) {
	eb.Subscriptions = append(eb.Subscriptions, subscript)
	// go func(blocks []func(int)) {
	// 	for v := range eb.Observable {
	// 		for _, sub := range blocks {
	// 			eb.wg.Add(1)
	// 			go func(value int) {
	// 				defer eb.wg.Done()
	// 				sub(value)
	// 			}(v)
	// 		}
	// 	}
	// }(eb.Subscriptions)
	go func() {
		for v := range eb.Observable {
			for _, sub := range eb.Subscriptions {
				sub(v)
			}
		}
	}()
	// 되는 코드 2
	// go func() {
	// 	for {
	// 		select {
	// 		case v, ok := <-eb.Observable:
	// 			if !ok {
	// 				return
	// 			}
	// 			for _, sub := range eb.Subscriptions {
	// 				eb.wg.Add(1)
	// 				go func(value int) {
	// 					defer eb.wg.Done()
	// 					sub(value)
	// 				}(v)
	// 			}
	// 		}
	// 	}
	// }()
}

func (eb *IntEventBus) Next(v int) {
	eb.Observable <- v
}

func (eb *IntEventBus) Dispose() {
	close(eb.Observable)
	eb.wg.Wait()
}
