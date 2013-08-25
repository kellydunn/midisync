package main

import (
	"fmt"
	"github.com/kellydunn/go-cmd"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	CLOCK       = 0xF8
	START       = 0xFA
	CONTINUE    = 0xFB
	STOP        = 0xFC
	MINUTE      = 60.0
	PPQN        = 24.0
	MICROSECOND = 1000000000
)

func main() {
	c := &cmd.Command{Name: "sync", Summary: "syncs the rc looper at the specified path to the specified bpm", Usage: "midisync sync /path/to/device bpm"}
	c.Run = sync
	cmd.Handle("sync", c)
	cmd.Serve()
}

func sync(command *cmd.Command, args ...string) {

	// Open the midi device
	fd, openErr := os.OpenFile(args[1], os.O_RDWR, 0666)
	if openErr != nil {
		fmt.Printf("Could not open the specified midi device")
		os.Exit(1)
	}

	bpm, _ := strconv.ParseFloat(args[2], 64)

	// Send START message
	_, startErr := fd.Write([]byte{START})
	if startErr != nil {
		fmt.Printf("Could not send Start message to midi device")
		os.Exit(1)
	}

	// Send CLOCK message immediately after for first downbeat
	_, clockErr := fd.Write([]byte{CLOCK})
	if clockErr != nil {
		fmt.Printf("Could not send Clock message to midi device")
		os.Exit(1)
	}

	// Send STOP message if user bails using SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go handleSigint(c, fd)

	// Determine interval in which to send midi signals
	interval := microsecondsPerPulse(bpm)

	// Sleep for the determined amount of time
	// Then write another CLOCK message
	for {
		time.Sleep(interval)
		fd.Write([]byte{CLOCK})
	}

}

func microsecondsPerPulse(bpm float64) time.Duration {
	return time.Duration((MINUTE * MICROSECOND) / (PPQN * bpm))
}

func handleSigint(c chan os.Signal, fd *os.File) {
	for sig := range c {
		if sig == syscall.SIGINT {
			fd.Write([]byte{STOP})
			os.Exit(0)
		}
	}
}
