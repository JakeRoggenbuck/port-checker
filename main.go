package main

import (
	"flag"
	"fmt"
	"github.com/cakturk/go-netstat/netstat"
	"time"
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

	var tcpall []netstat.SockTabEntry
	tcpall = append(tcpall, tabs6[:]...)
	tcpall = append(tcpall, tabs[:]...)
	return tcpall
}

func single_run(ports []netstat.SockTabEntry) {

	fmt.Printf("%12v %12v %12v %12v\n", "Local IP", "Local Port", "Proc Name", "Proc ID")
	fmt.Printf("%12v %12v %12v %12v\n", "--------", "----------", "---------", "-------")
	for _, p := range ports {
		fmt.Printf("%12v %12v", p.LocalAddr.IP, p.LocalAddr.Port)
		if p.Process != nil {
			fmt.Printf(" %12v %12v\n", p.Process.Name, p.Process.Pid)
		} else {
			fmt.Printf(" %12v %12v\n", " --- ", " --- ")
		}
	}
}

func clear_screen() {
	fmt.Println("\033[2J")
}

func main() {
	var seconds int
	flag.IntVar(&seconds, "seconds", 2, "update delay")

	var watch bool
	flag.BoolVar(&watch, "watch", false, "update view in watch mode")

	flag.Parse()

	seconds_dur := time.Duration(seconds)

	dot_count := 0
	if watch {
		for {
			tcp6Ports := getOpenTCP6Ports()
			clear_screen()
			single_run(tcp6Ports)

			// Print dots to show updates
			x := 0
			fmt.Print("\nRunning")
			for x < dot_count {
				fmt.Print(".")
				x += 1
			}
			if dot_count < 4 {
				dot_count += 1
			} else {
				dot_count = 0
			}

			time.Sleep(seconds_dur * time.Second)
		}
	} else {
		tcp6Ports := getOpenTCP6Ports()
		single_run(tcp6Ports)
	}
}
