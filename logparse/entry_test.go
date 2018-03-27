package logparse

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testingSourceIP         = "31.184.238.128"
	validLogEntryLineFormat = `%v - - [02/Jun/2015:17:00:12 -0700] "GET /logs/access.log HTTP/1.1" 200 2145998 "http://kmprograf.forumcircle.com" "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.146 Safari/537.36" "redlug.com"`
)

var testingEntry *LogEntry

func TestParseMalformedSourceIPFromEntry(t *testing.T) {
	malformedIP := "31.184.238.1285"
	entry := &LogEntry{
		fmt.Sprintf(validLogEntryLineFormat, malformedIP),
	}
	assert.Nil(t, entry.SourceIP())
}

func TestParseSourceIPFromMalformedEntry(t *testing.T) {
	entry := &LogEntry{
		fmt.Sprintf(`%v - [02/Jun/2015:17:00:12 -0700] "GET /logs/access.log HTTP/1.1" 200 2145998 "http://kmprograf.forumcircle.com" "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.146 Safari/537.36" "redlug.com"`, testingSourceIP),
	}
	assert.Nil(t, entry.SourceIP())
}

func TestParseSourceIPFromEntry(t *testing.T) {
	entry := &LogEntry{
		fmt.Sprintf(validLogEntryLineFormat, testingSourceIP),
	}
	fmt.Println(entry.logLine)
	assert.Equal(t, testingSourceIP, entry.SourceIP().String())
}

func TestMatchIP(t *testing.T) {
	entry := &LogEntry{
		fmt.Sprintf(validLogEntryLineFormat, testingSourceIP),
	}
	assert.True(t, entry.MatchIP(net.ParseIP(testingSourceIP)))
}

func TestMatchCIDR(t *testing.T) {
	entry := &LogEntry{
		fmt.Sprintf(validLogEntryLineFormat, testingSourceIP),
	}
	_, matchingCIDR, _ := net.ParseCIDR("31.184.238.0/24")
	_, nonMatchingCIDR, _ := net.ParseCIDR("31.184.237.0/24")
	assert.True(t, entry.MatchCIDR(matchingCIDR))
	assert.False(t, entry.MatchCIDR(nonMatchingCIDR))
}
