package style

import (
	"fmt"
	"os"
)

var enabled = isTTY()

type partType string

const (
	Text        partType = "text"
	Code        partType = "code"
	ComplexCode partType = "complex"
)

type part struct {
	value string
	kind  partType
}

type Style struct {
	parts   []part
	enabled bool
}

func isTTY() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

func New() Style {
	return Style{
		parts:   []part{},
		enabled: enabled,
	}
}

func (s Style) Enable() Style {
	s.enabled = true
	return s
}

func (s Style) Disable() Style {
	s.enabled = false
	return s
}

func (s Style) add(value string, kind partType) Style {
	newPart := part{
		value: value,
		kind:  kind,
	}
	newParts := append([]part{}, s.parts...)
	newParts = append(newParts, newPart)
	return Style{parts: newParts, enabled: s.enabled}
}

func (s Style) Bold() Style      { return s.add("1", Code) }
func (s Style) Dim() Style       { return s.add("2", Code) }
func (s Style) Italic() Style    { return s.add("3", Code) }
func (s Style) Underline() Style { return s.add("4", Code) }
func (s Style) Blink() Style     { return s.add("5", Code) }
func (s Style) Reverse() Style   { return s.add("7", Code) }
func (s Style) Strike() Style    { return s.add("9", Code) }

func (s Style) Black() Style   { return s.add("30", Code) }
func (s Style) Red() Style     { return s.add("31", Code) }
func (s Style) Green() Style   { return s.add("32", Code) }
func (s Style) Yellow() Style  { return s.add("33", Code) }
func (s Style) Blue() Style    { return s.add("34", Code) }
func (s Style) Magenta() Style { return s.add("35", Code) }
func (s Style) Cyan() Style    { return s.add("36", Code) }
func (s Style) White() Style   { return s.add("37", Code) }

func (s Style) BrightBlack() Style   { return s.add("90", Code) }
func (s Style) BrightRed() Style     { return s.add("91", Code) }
func (s Style) BrightGreen() Style   { return s.add("92", Code) }
func (s Style) BrightYellow() Style  { return s.add("93", Code) }
func (s Style) BrightBlue() Style    { return s.add("94", Code) }
func (s Style) BrightMagenta() Style { return s.add("95", Code) }
func (s Style) BrightCyan() Style    { return s.add("96", Code) }
func (s Style) BrightWhite() Style   { return s.add("97", Code) }

func (s Style) BgBlack() Style   { return s.add("40", Code) }
func (s Style) BgRed() Style     { return s.add("41", Code) }
func (s Style) BgGreen() Style   { return s.add("42", Code) }
func (s Style) BgYellow() Style  { return s.add("43", Code) }
func (s Style) BgBlue() Style    { return s.add("44", Code) }
func (s Style) BgMagenta() Style { return s.add("45", Code) }
func (s Style) BgCyan() Style    { return s.add("46", Code) }
func (s Style) BgWhite() Style   { return s.add("47", Code) }

func (s Style) BgBrightBlack() Style   { return s.add("100", Code) }
func (s Style) BgBrightRed() Style     { return s.add("101", Code) }
func (s Style) BgBrightGreen() Style   { return s.add("102", Code) }
func (s Style) BgBrightYellow() Style  { return s.add("103", Code) }
func (s Style) BgBrightBlue() Style    { return s.add("104", Code) }
func (s Style) BgBrightMagenta() Style { return s.add("105", Code) }
func (s Style) BgBrightCyan() Style    { return s.add("106", Code) }
func (s Style) BgBrightWhite() Style   { return s.add("107", Code) }

func (s Style) Fg256(n int) Style {
	return s.add(fmt.Sprintf("38;5;%d", n), ComplexCode)
}

func (s Style) Bg256(n int) Style {
	return s.add(fmt.Sprintf("48;5;%d", n), ComplexCode)
}

func (s Style) RGB(r, g, b int) Style {
	return s.add(fmt.Sprintf("38;2;%d;%d;%d", r, g, b), ComplexCode)
}

func (s Style) BgRGB(r, g, b int) Style {
	return s.add(fmt.Sprintf("48;2;%d;%d;%d", r, g, b), ComplexCode)
}

func (s Style) String(text string) Style {
	return s.add(text, Text)
}

func (s Style) Sprintf(format string, args ...any) Style {
	return s.add(fmt.Sprintf(format, args...), Text)
}

func (s Style) Reset() Style { return s.add("0", Code) }

func (s Style) Render() string {
	var out string
	finalStyle := s.Reset()

	for _, part := range finalStyle.parts {
		if part.kind == Text {
			out += part.value
		} else if part.kind == Code && finalStyle.enabled {
			out += sgr(part.value)
		} else if part.kind == ComplexCode && finalStyle.enabled {
			out += part.value
		}
	}

	return out
}

func Red() Style     { return New().Red() }
func Green() Style   { return New().Green() }
func Blue() Style    { return New().Blue() }
func Yellow() Style  { return New().Yellow() }
func Cyan() Style    { return New().Cyan() }
func Magenta() Style { return New().Magenta() }
func White() Style   { return New().White() }
