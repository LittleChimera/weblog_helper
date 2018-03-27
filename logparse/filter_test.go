package logparse

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleLogLines = []string{
	`31.184.238.128 - - [02/Jun/2015:17:00:12 -0700] "GET /logs/access.log HTTP/1.1" 200 2145998 "http://kmprograf.forumcircle.com" "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.146 Safari/537.36" "redlug.com"`,
	`157.55.39.180 - - [02/Jun/2015:17:00:46 -0700] "GET /Leaflets/2001/?M=D HTTP/1.1" 200 451 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)" "redlug.com"`,
	`151.80.31.151 - - [02/Jun/2015:17:04:32 -0700] "GET /paper2004/0401BNP.htm HTTP/1.1" 200 1354 "-" "Mozilla/5.0 (compatible; AhrefsBot/5.0; +http://ahrefs.com/robot/)" "redlug.com"`,
	`180.76.15.135 - - [02/Jun/2015:17:05:23 -0700] "GET /logs/access_140730.log HTTP/1.1" 200 979626 "-" "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)" "www.redlug.com"`,
	`180.76.15.137 - - [02/Jun/2015:17:05:28 -0700] "GET /logs/access_140730.log HTTP/1.1" 200 7849856 "-" "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)" "www.redlug.com"`,
	`93.79.202.178 - - [02/Jun/2015:17:05:41 -0700] "GET /logs/access_130930.log HTTP/1.1" 404 73 "http://kamagra.onlc.eu/1-generic-kamagra-jelly.html" "Mozilla/5.0 (Windows NT 6.1; rv:11.0) Gecko/20100101 Firefox/11.0" "redlug.com"`,
	`93.79.202.178 - - [02/Jun/2015:17:05:42 -0700] "GET /logs/access_130930.log HTTP/1.1" 404 73 "http://kamagra.onlc.eu/1-generic-kamagra-jelly.html" "Mozilla/5.0 (Windows NT 6.0; rv:17.0) Gecko/17.0 Firefox/17.0" "redlug.com"`,
}

var logBuffer io.Reader

func InitLogBuffer() {
	logBuffer = bytes.NewBufferString(strings.Join(sampleLogLines, "\n"))
}

func TestFilterLogsByIPWithoutResult(t *testing.T) {
	InitLogBuffer()
	out := bytes.NewBufferString("")
	FilteredByIP(logBuffer, net.ParseIP("5.5.5.5"), out)

	filteredContents, err := ioutil.ReadAll(out)
	assert.Nil(t, err)
	lines := strings.Split(string(filteredContents), "\n")

	assert.Equal(t, 1, len(lines))
	assert.True(t, len(lines[0]) == 0)
}

func TestFilterLogsByIP(t *testing.T) {
	InitLogBuffer()
	out := bytes.NewBufferString("")
	FilteredByIP(logBuffer, net.ParseIP("180.76.15.135"), out)

	filteredContents, err := ioutil.ReadAll(out)
	assert.Nil(t, err)
	lines := strings.Split(string(filteredContents), "\n")

	assert.Equal(t, 1, len(lines))
	assert.True(t, len(lines[0]) > 0)
}
