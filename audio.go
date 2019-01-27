package opro

import (
	"encoding/binary"
	"gordonklaus/portaudio"
	"net"
	"net/http"
	"sync"
)

// AudioServer contain output from audio server
type AudioServer struct {
	pc     net.PacketConn
	buffer []byte
	stream *portaudio.Stream
}

// AudioConn hold udp connection to the server
type AudioConn struct {
	conn     *net.UDPConn
	bufferIn []byte
}

const (
	sampleRate = 48000
	seconds    = 1.5
)

// RunAudio is starting audio
func RunAudio() (*portaudio.Stream, error) {
	portaudio.Initialize()
	// defer portaudio.Terminate()
	buffer := make([]float32, sampleRate*seconds)
	var wg sync.WaitGroup
	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, sampleRate*seconds, func(out []float32) {
		wg.Add(1)
		for i := 0; i < 1; i++ {
			go func() {
				defer wg.Done()
				resp, err := http.Get("http://192.168.1.9:8081/audio")
				chk(err)
				binary.Read(resp.Body, binary.BigEndian, &buffer)
				for i := range out {
					out[i] = buffer[i]
				}
			}()
		}
		wg.Wait()
	})
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func (pc *AudioServer) playAudio() error {
	if err := pc.stream.Start(); err != nil {
		return err
	}
	return nil
}

func (pc *AudioServer) stopAudio() error {
	if err := pc.stream.Stop(); err != nil {
		return err
	}
	return nil
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
