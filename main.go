package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("starting n0tify...")
	es := newEventServer()

	go listen(es)

	go es.startServer()

	addSlaves(es)

	// if we get a sigterm or sigint signals start shutdown process
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Printf("\nshutting down, got kill signal...\n")
}

func handleWorkspace(e event) {
	log.Println("running workspace slave...")
	notify(fmt.Sprintf("switched to workspace %s", e.eventParams[0]), "low", "Hypr")
	speak(fmt.Sprintf("workspace %s", e.eventParams[0]))
}

func handleActivewindow(e event) {
	log.Println("running activewindow slave...")
	notify(fmt.Sprintf("switched to window %s", e.eventParams[0]), "low", "Hypr")
}

func addSlaves(es *eventServer) {
	es.addSlave(workspace, handleWorkspace)
	es.addSlave(activewindow, handleActivewindow)
}
