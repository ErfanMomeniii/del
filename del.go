package del

import "sync"

type Listener struct {
	f func()
}

type Event struct {
	sync.Mutex
	name      string
	listeners []Listener
}

func NewEvent(name string, listeners ...Listener) Event {
	var listenerSlice []Listener

	for _, listener := range listeners {
		listenerSlice = append(listenerSlice, listener)
	}

	return Event{
		name:      name,
		listeners: listenerSlice,
	}
}

func NewListener(f func()) Listener {
	return Listener{f: f}
}

func (e *Event) AddListener(listeners ...Listener) {
	for _, listener := range listeners {
		e.listeners = append(e.listeners, listener)
	}
}

func (e *Event) Dispatch() {
	var (
		channel = make(chan interface{})
		wg      = new(sync.WaitGroup)
	)

	go func() {
		wg.Add(len(e.listeners))

		for _, l := range e.listeners {

			go func(f func()) {
				e.Lock()
				defer e.Unlock()

				f()

				wg.Done()
			}(l.f)

		}

		wg.Wait()

		channel <- "done"
	}()

	<-channel
}
