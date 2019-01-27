package opro

import (
	"net"
)

// Dial function is use to dial in audio server
// this return an audio connection
func Dial(address string) (*AudioConn, error) {
	raddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return nil, err
	}

	return &AudioConn{
		conn:     conn,
		bufferIn: make([]byte, 1024),
	}, nil

}

// SendCommand is a function for sending or write command data to the server
func (aux *AudioConn) SendCommand(cmd byte) (bool, error) {
	ok, err := aux.sendCommand(cmd)
	if !ok || err != nil {
		return false, err
	}
	return true, nil
}

// Close is for close audio client connection
func (aux *AudioConn) Close() error {
	if err := aux.conn.Close(); err != nil {
		return err
	}
	return nil
}

func (aux *AudioConn) sendCommand(v byte) (bool, error) {
	mode := byte(0x81)
	mask := byte(0x80)
	frame := []byte{mode, mask}
	sd := append(frame, v)
	_, err := aux.conn.Write(sd)
	if err != nil {
		return false, err
	}

	ok, err := aux.readStatus()
	if !ok || err != nil {
		return false, err
	}
	return true, nil

}

func (aux *AudioConn) readStatus() (bool, error) {
	_, err := aux.conn.Read(aux.bufferIn)
	if err != nil {
		return false, err
	}

	if aux.bufferIn[0] == StatusNotOK {
		return false, nil
	}
	return true, nil
}
