package log

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/ElementalMP4/loggerino/style"
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
		styledPrefix = style.New().Green().Bold().String(prefix).Render()
	case LevelInfo:
		styledPrefix = style.New().Blue().Bold().String(prefix).Render()
	case LevelWarn:
		styledPrefix = style.New().Yellow().Bold().String(prefix).Render()
	case LevelError:
		styledPrefix = style.New().Red().Bold().String(prefix).Render()
	case LevelDebug:
		styledPrefix = style.New().Magenta().Dim().String(prefix).Render()
	case LevelFatal:
		styledPrefix = style.New().BgRed().Bold().String(prefix).Render()
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

	if level == LevelFatal {
		os.Exit(1)
	}
}

// Text logs

func (l *Logger) Ok(source, msg string) {
	l.write(LevelOk, "  OK ", source, msg)
}

func (l *Logger) Info(source, msg string) {
	l.write(LevelInfo, " INFO", source, msg)
}

func (l *Logger) Warn(source, msg string) {
	l.write(LevelWarn, " WARN", source, msg)
}

func (l *Logger) Error(source, msg string) {
	l.write(LevelError, " ERR ", source, msg)
}

func (l *Logger) Debug(source, msg string) {
	l.write(LevelDebug, " DEBG", source, msg)
}

func (l *Logger) Fatal(source, msg string) {
	l.write(LevelFatal, "FATAL", source, msg)
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

func (l *Logger) Fatalf(source, format string, v ...any) {
	l.Fatal(source, fmt.Sprintf(format, v...))
}

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = 200
	}
	n, err := rw.ResponseWriter.Write(b)
	rw.size += n
	return n, err
}

func (l *Logger) LoggingMiddleware(info RequestInfo) {
	statusStr := fmt.Sprintf("%3d", info.Status)
	var styledStatus string

	switch {
	case info.Status >= 200 && info.Status < 300:
		styledStatus = style.New().Green().Bold().String(statusStr).Render()
	case info.Status >= 300 && info.Status < 400:
		styledStatus = style.New().Cyan().Bold().String(statusStr).Render()
	case info.Status >= 400 && info.Status < 500:
		styledStatus = style.New().Yellow().Bold().String(statusStr).Render()
	default:
		styledStatus = style.New().Red().Bold().String(statusStr).Render()
	}

	var styledMethod string
	switch info.Method {
	case "GET":
		styledMethod = style.New().Blue().Bold().String(info.Method).Render()
	case "POST":
		styledMethod = style.New().Green().Bold().String(info.Method).Render()
	case "PUT":
		styledMethod = style.New().Yellow().Bold().String(info.Method).Render()
	case "DELETE":
		styledMethod = style.New().Red().Bold().String(info.Method).Render()
	case "PATCH":
		styledMethod = style.New().Magenta().Bold().String(info.Method).Render()
	default:
		styledMethod = info.Method
	}

	latency := info.Latency.String()

	msg := fmt.Sprintf("%s %s %s (%s)", styledMethod, info.Path, styledStatus, latency)

	level := LevelInfo
	if info.Status >= 400 {
		level = LevelError
	}
	l.write(level, styledStatus, info.IP, msg)
}
