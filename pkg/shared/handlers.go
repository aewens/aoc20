package shared

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Cleanup() {
	r := recover()
	if r != nil {
		log.Println("[!]", r)
	}
}

func HandleSigterm() {
	sigterm := make(chan os.Signal)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigterm
		Cleanup()
	}()
}
