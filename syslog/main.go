package main

import (
	"fmt"
	"log"
	"log/syslog"
	"time"
)

func main() {
	l3, err := syslog.Dial("udp", "3.95.59.242:30514", syslog.LOG_ERR, "GoExample 1") // connection to a log daemon
	defer l3.Close()
	if err != nil {
		log.Fatal("error")
	}

	l4, err := syslog.Dial("udp", "54.196.12.126:30514", syslog.LOG_ERR, "GoExample 2") // connection to a log daemon
	defer l4.Close()
	if err != nil {
		log.Fatal("error")
	}

	l5, err := syslog.Dial("udp", "3.89.79.171:30514", syslog.LOG_ERR, "GoExample 3") // connection to a log daemon
	defer l5.Close()
	if err != nil {
		log.Fatal("error")
	}

	for true {
		l3.Info(fmt.Sprintf("Go dh 1 %d", time.Now().UnixNano()))
		l4.Info(fmt.Sprintf("Go dh 2 %d", time.Now().UnixNano()))
		l5.Info(fmt.Sprintf("Go dh 2 %d", time.Now().UnixNano()))
	}

}
