package sender

import (
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p-gossip/packet"
	"github.com/kodykantor/p2p-gossip/udp/peer"
)

type Sender struct {
	peer *peer.Peer
}

func init() {
}

func New(peer *peer.Peer) *Sender {
	//TODO check if peer is nil
	newSender := new(Sender)
	newSender.peer = peer
	return newSender
}

const (
	senderAddr   = "localhost:8090"
	receiverAddr = "localhost:8080"
)

func (s *Sender) Send(ch chan packet.Packet) error {
	logrus.Debugln("Starting to send packet...")
	//	senderAddr := "localhost:" + strconv.Itoa(s.peer.GetPort()) //TODO change this
	//	receiverAddr := "localhost:" + strconv.Itoa(s.peer.GetPort())

	logrus.Printf("Client's listen address should be: %v", receiverAddr)
	la, err := net.ResolveUDPAddr("udp", receiverAddr)
	if err != nil {
		return fmt.Errorf("Error resolving the listener's address: %v", err)
	}

	logrus.Debugln("Starting packet listener for sender")
	sendConn, err := net.ListenPacket("udp", senderAddr)
	if err != nil {
		return fmt.Errorf("Error getting send connection: %v", err)
	}
	defer sendConn.Close()

	logrus.Debugln("Reading from channel to send.")
	pack := <-ch //read a packet to send from the channel
	logrus.Debugln("Received packet from channel to senD.")
	buf := pack.GetBuffer()

	count, err := sendConn.(*net.UDPConn).WriteToUDP(buf, la)
	if err != nil {
		return fmt.Errorf("Error writing packet to UDP: %v", err)
	}
	logrus.Debugln("Sent ", count, "bytes.")
	logrus.Debugln("Sender sent packet!")

	return nil
}
