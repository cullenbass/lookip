package main

import (
	"net"
	"regexp"
	"strconv"
)

type LogData struct {
	Ip         net.IP
	StatusCode uint32
}

func parse(line string) *LogData {
	// Break out the standard Apache log vars into an array-
	// [1] - IP address
	// [8] - Status code
	re := regexp.MustCompile("^([0-9.]+)\\s([\\w.-]+)\\s([\\w.-]+)\\s\\[([^\\]]+)\\]\\s\"(\\w+)\\s(\\S+)\\s([^\"]+)\"\\s([0-9]+)\\s([0-9]+)\\s\"([^\"]+)\"\\s\"([^\"]+)\"")
	dat := re.FindStringSubmatch(line)
	// Grab the HTTP status code
	status, _ := strconv.Atoi(dat[8])

	return &LogData{
		Ip:         net.ParseIP(dat[1]),
		StatusCode: uint32(status),
	}
}
