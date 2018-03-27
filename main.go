package main

import (
	"flag"
	"io"
	"net"
	"net/http"
	"os"

	"./logparse"
)

const (
	logFileLocation = "https://s3.amazonaws.com/syseng-challenge/public_access.log.txt"
)

func main() {
	ipFlag := flag.String("ip", "", "")
	flag.Parse()

	if address := net.ParseIP(*ipFlag); address != nil {
		logReader := logReader()
		defer logReader.Close()
		logparse.FilteredByIP(logReader, address, os.Stdout)
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
