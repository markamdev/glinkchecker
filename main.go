package main

import (
	"fmt"
	"glinkchecker/checker"
	"glinkchecker/gateway"
	"os"
	"os/signal"

	"github.com/namsral/flag"
)

func main() {
	fmt.Println("Link stability checker")
	interval := flag.Int("interval", 30, "Interval between regular checks (in [s])")
	address := flag.String("address", "", "Address to be used for link chceking (given as IP)")
	port := flag.String("port", "80", "Port number for connection checking. By default set to 80 (http without SSL/TLS)")

	flag.Parse()

	if len(*address) == 0 {
		gw, err := gateway.DetectDefaultGateway()
		if err != nil {
			fmt.Println("Unable to detect gateway address:", err.Error())
			os.Exit(1)
		}
		address = &gw
	}

	checkerAddress := *address + ":" + *port

	fmt.Println("Checking connection to", checkerAddress, "with", *interval, "[s] interval")

	// wait for sigint
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	chk := checker.CreateChecker(checkerAddress, *interval)
	chk.Start()
	fmt.Println("Launched!")

	<-sigChan

	chk.Stop()
	fmt.Println("Connection checker disabled")
}
