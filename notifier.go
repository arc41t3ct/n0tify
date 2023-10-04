package main

import (
	"log"
	"os/exec"
)

// notify accepts a message and a level (low, normal, critical)
// t is the title
func notify(m string, l, t string) {
	log.Println("going to notify:", m)
	cmd := exec.Command(
		"notify-send",
		"-t", "800",
		"-u", l,
		t,
		m)
	log.Println(cmd.String())
	_, cErr := cmd.Output()
	if cErr != nil {
		log.Fatal(cErr)
	}
}
