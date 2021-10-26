package checker

import (
	"net"
	"time"
)

type HttpChecker interface {
	Start()
	Stop()
}

func CreateChecker(address string, interval int) HttpChecker {
	dlr := net.Dialer{
		Timeout:   time.Second * 5,
		DualStack: false,
	}
	if interval < 10 {
		interval = 10
	}
	return &simpleChecker{endpointAddress: address, normalInterval: interval, failureInterval: 10, tcpCheck: dlr}
}
