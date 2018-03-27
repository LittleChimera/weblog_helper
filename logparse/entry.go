package logparse

import (
	"net"
	"strings"
)

type LogEntry struct {
	logLine string
}

func (entry *LogEntry) SourceIP() net.IP {
	splitLogLine := strings.Split(entry.logLine, " - - ")
	if len(splitLogLine) < 2 {
		return nil
	}

	ipAddress := net.ParseIP(splitLogLine[0])
	if ipAddress != nil {
		return ipAddress
	}

	return nil
}

func (entry *LogEntry) MatchIP(address net.IP) bool {
	return address.Equal(entry.SourceIP())
}

func (entry *LogEntry) MatchCIDR(mask *net.IPNet) bool {
	return mask.Contains(entry.SourceIP())
}
