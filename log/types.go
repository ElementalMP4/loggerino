package log

import (
	"io"
	"sync"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelOk
	LevelWarn
	LevelError
	LevelFatal
)

type Logger struct {
	out        io.Writer
	err        io.Writer
	file       io.Writer
	minLevel   Level
	timeFormat string
	mu         sync.Mutex
}
