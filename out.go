package opro

import (
	"fmt"
	"net"

	"github.com/google/logger"
)

// Listen running audio server
func Listen(addr string) (*AudioServer, error) {
	lis, err := net.ListenPacket("udp", addr)
	if err != nil {
		logger.Fatalf("error resolve addr : %v\n", err)
		return nil, err
	}
	stream, err := RunAudio()
	if err != nil {
		return nil, err
	}
	return &AudioServer{
		buffer: make([]byte, 1024),
		pc:     lis,
		stream: stream,
	}, nil
}

// ReadData is read data from client
func (pc *AudioServer) ReadData() {
	for {
		_, addr, err := pc.pc.ReadFrom(pc.buffer)
		if err != nil {
			logger.Infof("error read data: %v", err)
		}
		ok, err := pc.readHeader(addr)
		if err != nil {
			logger.Errorf("error header : %v", err)
		}
		if ok {
			pc.process(addr)
		}

	}
}

func (pc *AudioServer) process(addr net.Addr) {
	res := pc.buffer[2]
	switch res {
	case PlayAudio:
		fmt.Println("play audio")
		pc.playAudio()
		pc.sendStatus(StatusOK, addr)
	case StopAudio:
		fmt.Println("stop audio")
		pc.stopAudio()
		pc.sendStatus(StatusOK, addr)
	default:
		fmt.Println("command error")
		pc.sendStatus(StatusNotOK, addr)
	}
}

func (pc *AudioServer) sendStatus(status byte, addr net.Addr) error {
	msg := []byte{status}
	_, err := pc.pc.WriteTo(msg, addr)
	if err != nil {
		return err
	}
	return nil
}

// Close is to close audio server
func (pc *AudioServer) Close() error {
	if err := pc.pc.Close(); err != nil {
		return err
	}
	return nil
}

func (pc *AudioServer) readHeader(addr net.Addr) (bool, error) {
	err := checkHeader(pc.buffer)
	if err != nil {
		if err == errMaskInvalid || err == errModeInvalid {
			err := pc.sendStatus(StatusNotOK, addr)
			if err != nil {
				return false, err
			}
		}
		return false, err
	}
	return true, nil
}

func checkHeader(r []byte) error {
	if len(r) == 0 {
		return errEmptyByte
	}
	if r[0] != 0x81 {
		return errModeInvalid
	}
	if (r[1] & 0x80) != 0x80 {
		return errMaskInvalid
	}
	return nil
}
