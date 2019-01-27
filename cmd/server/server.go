package main

import (
	"flag"
	"os"

	"github.com/fn-code/opro"
	"github.com/google/logger"
)

var (
	port    string
	Lg      *logger.Logger
	logFile string = "./logger/log.log"
)

func main() {

	flag.StringVar(&port, "port", ":5005", "port you want to use")
	flag.Parse()

	fl, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("failed to open file: %v", err)
	}
	defer fl.Close()
	Lg = logger.Init("AudioServer", true, true, fl)
	defer Lg.Close()
	logger.Infof("server running on port%s\n", port)

	lis, err := opro.Listen(port)
	if err != nil {
		logger.Error(err)
	}
	defer lis.Close()
	lis.ReadData()

}
