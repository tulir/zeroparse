package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := &zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
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
