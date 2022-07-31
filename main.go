package main

import (
	"fmt"
	"github.com/cakturk/go-netstat/netstat"
)

func getOpenTCP6Ports() []netstat.SockTabEntry {
	tabs6, err := netstat.TCP6Socks(func(s *netstat.SockTabEntry) bool {
		return s.State == netstat.Listen
	})
	if err != nil {
		fmt.Println(err)
	}

	tabs, err := netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
		return s.State == netstat.Listen
	})
	if err != nil {
		fmt.Println(err)
	}

	var tcpall [] netstat.SockTabEntry
	tcpall = append(tcpall, tabs6[:]...)
	tcpall = append(tcpall, tabs[:]...)
	return tcpall
}

func main() {
	tcp6Ports := getOpenTCP6Ports()

	fmt.Printf("%12v %12v %12v %12v\n", "Local IP", "Local Port", "Proc Name", "Proc ID")
	for _, p := range tcp6Ports {
		fmt.Printf("%12v %12v", p.LocalAddr.IP,  p.LocalAddr.Port)
		if p.Process != nil {
			fmt.Printf(" %12v %12v\n", p.Process.Name, p.Process.Pid)
		} else {
			fmt.Printf(" %12v %12v\n", " --- ", " --- ")
		}
	}
}
