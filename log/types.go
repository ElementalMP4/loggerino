package log

import (
	"io"
	"sync"
)

type Level int

const (
	LevelDebug Level = iota
	LevelOk
	LevelInfo
	LevelWarn
	LevelError
)

type Logger struct {
	out        io.Writer
	err        io.Writer
	file       io.Writer
	minLevel   Level
	timeFormat string
	mu         sync.Mutex
}
