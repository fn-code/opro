package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fn-code/opro"
	"github.com/google/logger"
)

const (
	address = "0.0.0.0:5005"
)

var (
	// Lg is use for define logger
	Lg *logger.Logger
)

func main() {
	fl, err := os.OpenFile("./logger/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("failed to open file: %v\n", err)
	}

	// Set logger verbose to false if don't want logger to show in stdout
	Lg = logger.Init("AudioClient", true, true, fl)
	defer Lg.Close()

	conn, err := opro.Dial(address)
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()
	r := bufio.NewScanner(os.Stdin)
	for r.Scan() {
		s := r.Text()
		switch s {
		case "1":
			ok, err := conn.SendCommand(opro.PlayAudio)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(ok)
		case "2":
			ok, err := conn.SendCommand(opro.StopAudio)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(ok)
		case "3":
			ok, err := conn.SendCommand(byte(0x44))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(ok)
		}
	}

}
