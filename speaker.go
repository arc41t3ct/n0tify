package main

import (
	"log"
	"os/exec"
)

func speak(t string) {
	log.Println("going to say:", t)
	cmd := exec.Command(
		"espeak",
		"-ven-us+f2",
		t,
		"-g02ms",
		"-a180",
		"-p99",
		"-s210")
	log.Println(cmd.String())
	_, cErr := cmd.Output()
	if cErr != nil {
		log.Fatal(cErr)
	}
}
