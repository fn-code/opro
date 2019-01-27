package opro

import "errors"

var (
	errModeInvalid   = errors.New("error mode is invalid")
	errMaskInvalid   = errors.New("error mask is invalid")
	errEmptyByte     = errors.New("error receive empty data")
	errHeaderInvalid = errors.New("header invalid")
)

const (
	// StatusOK is status ok
	StatusOK = byte(0x80)
	// StatusNotOK is status is not ok
	StatusNotOK = byte(0x81)
	// PlayAudio is command to play audio
	PlayAudio = byte(0x33)
	// StopAudio is command to stop audio
	StopAudio = byte(0x34)
)
