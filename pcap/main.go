package main

import (
	"log"
	"log/syslog"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	handle, err := pcap.OpenOffline("/Volumes/Data/graylog-testdata/tempdir/dump-udp-514.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	l3, err := syslog.Dial("udp", "3.95.59.242:30514", syslog.LOG_ERR, "GoPKG") // connection to a log daemon
	defer l3.Close()
	if err != nil {
		log.Fatal("error")
	}

	l5, err := syslog.Dial("udp", "3.89.79.171:30514", syslog.LOG_ERR, "GoExample 3") // connection to a log daemon
	defer l5.Close()
	if err != nil {
		log.Fatal("error")
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	r := 0
	for packet := range packetSource.Packets() {
		r++

		time.Sleep(1 * time.Second / 7000)

		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {

			udp, _ := udpLayer.(*layers.UDP)
			if r%2 == 0 {
				l5.Info(string(udp.Payload))
			} else {
				l3.Info(string(udp.Payload))
			}
		}
	}
}
