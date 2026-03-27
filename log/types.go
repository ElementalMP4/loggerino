package log

import (
	"io"
	"sync"
	"time"
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

type RequestInfo struct {
	Method  string
	Path    string
	Status  int
	IP      string
	Latency time.Duration
}
