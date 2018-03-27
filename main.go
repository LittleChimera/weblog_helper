package main

import (
	"flag"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/lukadante/weblog_helper/logparse"
)

const (
	logFileLocation = "https://s3.amazonaws.com/syseng-challenge/public_access.log.txt"
)

func main() {
	ipFlag := flag.String("ip", "", "")
	flag.Parse()

	logReader := logReader()
	defer logReader.Close()

	if address := net.ParseIP(*ipFlag); address != nil {
		logparse.FilteredByIP(logReader, address, os.Stdout)
	} else if _, mask, err := net.ParseCIDR(*ipFlag); err == nil {
		logparse.FilteredByCIDR(logReader, mask, os.Stdout)
	} else {
		panic("Invalid or missing --ip argument")
	}
}

func logReader() io.ReadCloser {
	logFileResponse, err := http.Get(logFileLocation)
	if err != nil {
		panic(err)
	}
	return logFileResponse.Body
}
