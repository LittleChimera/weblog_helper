package logparse

import (
	"bufio"
	"io"
	"net"
)

func FilteredByIP(reader io.Reader, address net.IP, out io.ReadWriter) {
	bufferedReader := bufio.NewReader(reader)

	for {
		logLine, _, err := bufferedReader.ReadLine()

		if err == io.EOF {
			return
		}
		entry := &LogEntry{string(logLine)}
		if entry.MatchIP(address) {
			out.Write(logLine)
		}
	}
}
