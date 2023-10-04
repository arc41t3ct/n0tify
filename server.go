package main

import (
	"log"
	"strings"
	"sync"
)

type eventServer struct {
	eventChan chan event
	events    map[int]event
	slaves    map[hyprEvent][]slave

	mu sync.Mutex
}

func newEventServer() *eventServer {
	return &eventServer{
		eventChan: make(chan event),
		events:    map[int]event{},
		slaves:    map[hyprEvent][]slave{},
	}
}

func (e *eventServer) dispatch(ev string) {
	log.Println("dispatching event...")
	var t hyprEvent
	var p []string
	split := strings.Split(ev, ">>")
	if len(split) < 2 {
		return
	}
	switch split[0] {
	case "openlayer":
		t = openlayer
		break
	case "closelayer":
		t = closelayer
		break
	case "activewindow":
		t = activewindow
		break
	case "activewindowv2":
		t = activewindowv2
		break
	case "workspace":
		t = workspace
		break
	case "destroyworkspace":
		t = destroyworkspace
		break
	default:
		t = unknown
		break
	}
	p = strings.Split(split[1], ",")
	ne := newEvent(t, p, ev)
	e.eventChan <- ne
}

func (e *eventServer) sameAsPrevious(newEv event) bool {
	e.mu.Lock()
	if e.events[len(e.events)-1].raw == newEv.raw {
		return true
	}
	e.mu.Unlock()
	return false
}

// we check wether the same event was already added if so we skip
// since hypr could send activewindow a bunch of times without leaving
func (e *eventServer) storeEvent(ev event) {
	log.Println("storing event...")
	if !e.sameAsPrevious(ev) {
		e.mu.Lock()
		e.events[len(e.events)] = ev
		e.mu.Unlock()
		log.Println("event stored...")
	}
}

func (e *eventServer) dispatchConsumers(ev event) {
	e.mu.Lock()
	if _, exists := e.slaves[ev.eventType]; exists {
		for _, c := range e.slaves[ev.eventType] {
			go c.eventWorker(ev)
		}
	}
	e.mu.Unlock()
}

func (e *eventServer) startServer() {
	log.Println("starting server...")
	for {
		select {
		case ev := <-e.eventChan:
			log.Println("event received:", ev.string())
			e.storeEvent(ev)
			e.dispatchConsumers(ev)
		}
	}
}

func (e *eventServer) addSlave(et hyprEvent, fn slaveWorker) {
	log.Println("adding slave...")
	s := slave{
		eventType:   et,
		eventWorker: fn,
	}
	e.mu.Lock()
	if _, exists := e.slaves[et]; !exists {
		e.slaves[et] = []slave{}
	}
	e.slaves[et] = append(e.slaves[et], s)
	e.mu.Unlock()
	log.Println("slave added for hyprEvent:", et)
}
