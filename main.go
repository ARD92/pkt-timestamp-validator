package main
import (
	"fmt"
	//"net"
	"github.com/google/gopacket"
	//"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/akamensky/argparse"
	"os"
	"time"
	"log"
)

var (
    snapshotLen int32  = 1024
    promiscuous bool   = false
    err         error
    timeout     time.Duration = 5 * time.Second
    handle      *pcap.Handle
)

func printNanoTS(packet gopacket.Packet) {
	timestamp := packet.Metadata().Timestamp
	nanoseconds := timestamp.UnixNano()
	fmt.Println("Timestamp:", nanoseconds)
}

func main() {
	logs, logerr := os.OpenFile("ts-validate.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if logerr != nil {
		log.Fatalf("Error opening file: %v", logerr)
	}
	defer logs.Close()
	log.SetOutput(logs)
	parser := argparse.NewParser("Required-args", "\n============\npacket-timestamp-validator\n============")
	device := parser.String("i", "intf", &argparse.Options{Required: true, Help: "interface to bind to "})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	} else {
		handle, err = pcap.OpenLive(*device, snapshotLen, promiscuous, timeout)
		if err != nil { log.Fatal(err) }
		defer handle.Close()
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			printNanoTS(packet)
		}
	}
}
