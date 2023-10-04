package main

import "log"

type slaveWorker func(ev event)

type slave struct {
	eventType   hyprEvent
	eventWorker slaveWorker
}

func (s *slave) consume(ev event) {
	log.Println("consuming:", ev.string())
	s.eventWorker(ev)
}
