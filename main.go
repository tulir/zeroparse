package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := &zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02T15:04:05.999Z07:00"}
	// Ignore SIGINT/SIGTERM, the process will exit when stdin is closed.
	// If we don't do this, then pressing ctrl+c will make zeroparse exit
	// without reading whatever logs are produced by the program during shutdown.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		if line, err := reader.ReadBytes('\n'); err != nil {
			// Don't log EOFs as that would just create log unnecessary noise when finishing reading logs
			if !errors.Is(err, io.EOF) {
				_, _ = fmt.Fprintf(os.Stderr, "Failed to read line: %v", err)
			}
			return
		} else if !json.Valid(line) {
			// Pass through non-JSON lines as-is
			_, _ = os.Stdout.Write(line)
		} else if _, err = writer.Write(line); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to write line: %v", err)
			return
		}
	}
}
