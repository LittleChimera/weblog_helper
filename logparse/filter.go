package logparse

import (
	"bufio"
	"io"
	"net"
)

func FilteredByIP(reader io.Reader, address net.IP, out io.ReadWriter) {
	filterAndWrite(reader, func(entry *LogEntry) bool {
		return entry.MatchIP(address)
	}, out)
}

func FilteredByCIDR(reader io.Reader, mask *net.IPNet, out io.ReadWriter) {
	filterAndWrite(reader, func(entry *LogEntry) bool {
		return entry.MatchCIDR(mask)
	}, out)
}

func filterAndWrite(reader io.Reader, filter func(entry *LogEntry) bool, out io.ReadWriter) {
	bufferedReader := bufio.NewReader(reader)

	for {
		logLine, _, err := bufferedReader.ReadLine()

		if err == io.EOF {
			return
		}
		entry := &LogEntry{string(logLine)}
		if filter(entry) {
			out.Write(logLine)
		}
	}
}
