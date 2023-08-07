package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Notify() {
	done := make(chan bool)
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal)
	go func() {
		for sig := range osSignal {
			switch sig {
			case syscall.SIGINT:
				fmt.Printf("SIGINT Signal %s received!\n", sig)
				done <- true
			case syscall.SIGTERM:
				fmt.Printf("mcSIGTERM Signal %s received!\n", sig)
				fmt.Println("in notify file..... sigterm")
			case syscall.SIGHUP:
				fmt.Printf("SIGHUP Signal %s received!\n", sig)
			case syscall.SIGPIPE:
				fmt.Printf("SIGPIPE Signal %s received!\n", sig)
			case syscall.SIGABRT:
				fmt.Printf("SIGABRT Signal %s received!\n", sig)
			case syscall.SIGURG:

			default:
				fmt.Printf("Unhandled Signal %s received!\n", sig)
			}
		}
	}()
	<-done

}
