package log

import (
	"fmt"
	"io"
	"loggerino/style"
	"os"
	"regexp"
	"time"
)

func New() *Logger {
	return &Logger{
		out:        os.Stdout,
		err:        os.Stderr,
		minLevel:   LevelDebug,
		timeFormat: "02-01-2006 15:04:05.000",
	}
}

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

func (l *Logger) SetLevel(level Level)       { l.minLevel = level }
func (l *Logger) SetOutput(w io.Writer)      { l.out = w }
func (l *Logger) SetErrorOutput(w io.Writer) { l.err = w }
func (l *Logger) SetTimeFormat(f string)     { l.timeFormat = f }

func (l *Logger) SetFile(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	l.file = f
	return nil
}

func (l *Logger) timestamp() string {
	return time.Now().Format(l.timeFormat)
}

func (l *Logger) write(level Level, prefix, source, msg string) {
	if level < l.minLevel {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	var styledPrefix string
	switch level {
	case LevelOk:
		styledPrefix = style.Green().Bold().String(prefix).Render()
	case LevelInfo:
		styledPrefix = style.Blue().Bold().String(prefix).Render()
	case LevelWarn:
		styledPrefix = style.Yellow().Bold().String(prefix).Render()
	case LevelError:
		styledPrefix = style.Red().Bold().String(prefix).Render()
	case LevelDebug:
		styledPrefix = style.Magenta().Dim().String(prefix).Render()
	default:
		styledPrefix = prefix
	}

	timestampStyled := style.New().BrightBlack().String(l.timestamp()).Reset().Render()
	sourcePadded := fmt.Sprintf("%15s", source)
	sourceStyled := style.New().Bold().Magenta().String(sourcePadded).Render()

	line := fmt.Sprintf("%s  %s  %s: %s\n", timestampStyled, styledPrefix, sourceStyled, msg)

	if level >= LevelError {
		l.err.Write([]byte(line))
	} else {
		l.out.Write([]byte(line))
	}

	if l.file != nil {
		plainLine := stripANSI(line)
		l.file.Write([]byte(plainLine))
	}
}

// Text logs

func (l *Logger) Ok(source, msg string) {
	l.write(LevelOk, " OK ", source, msg)
}

func (l *Logger) Info(source, msg string) {
	l.write(LevelInfo, "INFO", source, msg)
}

func (l *Logger) Warn(source, msg string) {
	l.write(LevelWarn, "WARN", source, msg)
}

func (l *Logger) Error(source, msg string) {
	l.write(LevelError, "FAIL", source, msg)
}

func (l *Logger) Debug(source, msg string) {
	l.write(LevelDebug, "DEBG", source, msg)
}

// With format

func (l *Logger) Okf(source, format string, v ...any) {
	l.Ok(source, fmt.Sprintf(format, v...))
}

func (l *Logger) Infof(source, format string, v ...any) {
	l.Info(source, fmt.Sprintf(format, v...))
}

func (l *Logger) Warnf(source, format string, v ...any) {
	l.Warn(source, fmt.Sprintf(format, v...))
}

func (l *Logger) Errorf(source, format string, v ...any) {
	l.Error(source, fmt.Sprintf(format, v...))
}

func (l *Logger) Debugf(source, format string, v ...any) {
	l.Debug(source, fmt.Sprintf(format, v...))
}
