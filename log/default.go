package log

import "net/http"

var std = New()

// Default text logger

func Ok(source, msg string)    { std.Ok(source, msg) }
func Info(source, msg string)  { std.Info(source, msg) }
func Warn(source, msg string)  { std.Warn(source, msg) }
func Error(source, msg string) { std.Error(source, msg) }
func Debug(source, msg string) { std.Debug(source, msg) }
func Fatal(source, msg string) { std.Fatal(source, msg) }

// Default format logger

func Okf(source, f string, v ...any)    { std.Okf(source, f, v...) }
func Infof(source, f string, v ...any)  { std.Infof(source, f, v...) }
func Warnf(source, f string, v ...any)  { std.Warnf(source, f, v...) }
func Errorf(source, f string, v ...any) { std.Errorf(source, f, v...) }
func Debugf(source, f string, v ...any) { std.Debugf(source, f, v...) }
func Fatalf(source, f string, v ...any) { std.Fatalf(source, f, v...) }

func Middleware(next http.Handler) http.Handler { return std.Middleware(next) }
