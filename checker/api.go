package checker

type HttpChecker interface {
	Start()
	Stop()
}

func CreateChecker(address string, interval int) HttpChecker {
	return &simpleChecker{endpointAddress: address, normalInterval: interval, failureInterval: 2}
}
