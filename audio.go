package opro

import (
	"net"
)

// AudioServer contain output from audio server
type AudioServer struct {
	pc     net.PacketConn
	buffer []byte
}

// AudioConn hold udp connection to the server
type AudioConn struct {
	conn     *net.UDPConn
	bufferIn []byte
}
