package main

import (
	"fmt"
	"strings"
)

type hyprEvent int

const (
	openlayer hyprEvent = iota
	closelayer
	activewindow
	activewindowv2
	workspace
	destroyworkspace
	unknown
)

type event struct {
	eventType   hyprEvent
	eventParams []string
	raw         string
}

func (e event) string() string {
	var n string
	switch e.eventType {
	case openlayer:
		n = "openlayer"
	case closelayer:
		n = "closelayer"
	case activewindow:
		n = "activewindow"
	case activewindowv2:
		n = "activewindowv2"
	case workspace:
		n = "workspace"
	case destroyworkspace:
		n = "destroyworkspace"
	case unknown:
		n = "unknown"
	}
	return fmt.Sprintf("%s, params: %s, raw: %s", n, strings.Join(e.eventParams, ", "), e.raw)
}

func newEvent(t hyprEvent, p []string, r string) event {
	return event{
		eventType:   t,
		eventParams: p,
		raw:         r,
	}
}
