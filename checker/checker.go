package checker

import (
	"fmt"
	"net"
	"time"
)

type simpleChecker struct {
	normalInterval  int
	failureInterval int
	endpointAddress string
	messages        chan bool
	tcpCheck        net.Dialer
}

func (sc *simpleChecker) Start() {
	fmt.Println("Stating checker at:", time.Now())
	sc.messages = make(chan bool)
	go sc.checkerProcess()
}

func (sc *simpleChecker) Stop() {
	fmt.Println("Stopping checker at:", time.Now())
	sc.messages <- true
}

func (sc *simpleChecker) checkerProcess() {
	currentInterval := sc.normalInterval
	for {
		select {
		case <-sc.messages:
			// stopchecking
			return
		case <-time.After(time.Second * time.Duration(currentInterval)):
			if sc.isAlive() {
				// do not register this event to minimize output size
				currentInterval = sc.normalInterval
			} else {
				// register failure with time
				sc.registerFailure()
				currentInterval = sc.failureInterval
			}
			fmt.Println("Timeout!")
		}
	}
}

func (sc *simpleChecker) isAlive() bool {
	conn, err := sc.tcpCheck.Dial("tcp", sc.endpointAddress)
	if err != nil {
		return false
	}

	conn.Close()
	return true
}

func (sc *simpleChecker) registerFailure() {
	fmt.Println("Connection failed at:", time.Now())
}
