package style

import "fmt"

func FgBlack() string   { return sgr("30") }
func FgRed() string     { return sgr("31") }
func FgGreen() string   { return sgr("32") }
func FgYellow() string  { return sgr("33") }
func FgBlue() string    { return sgr("34") }
func FgMagenta() string { return sgr("35") }
func FgCyan() string    { return sgr("36") }
func FgWhite() string   { return sgr("37") }
func FgDefault() string { return sgr("39") }

func FgBrightBlack() string   { return sgr("90") }
func FgBrightRed() string     { return sgr("91") }
func FgBrightGreen() string   { return sgr("92") }
func FgBrightYellow() string  { return sgr("93") }
func FgBrightBlue() string    { return sgr("94") }
func FgBrightMagenta() string { return sgr("95") }
func FgBrightCyan() string    { return sgr("96") }
func FgBrightWhite() string   { return sgr("97") }

func BgBlack() string   { return sgr("40") }
func BgRed() string     { return sgr("41") }
func BgGreen() string   { return sgr("42") }
func BgYellow() string  { return sgr("43") }
func BgBlue() string    { return sgr("44") }
func BgMagenta() string { return sgr("45") }
func BgCyan() string    { return sgr("46") }
func BgWhite() string   { return sgr("47") }
func BgDefault() string { return sgr("49") }

func BgBrightBlack() string   { return sgr("100") }
func BgBrightRed() string     { return sgr("101") }
func BgBrightGreen() string   { return sgr("102") }
func BgBrightYellow() string  { return sgr("103") }
func BgBrightBlue() string    { return sgr("104") }
func BgBrightMagenta() string { return sgr("105") }
func BgBrightCyan() string    { return sgr("106") }
func BgBrightWhite() string   { return sgr("107") }

func Fg(n int) string { return sgr(fmt.Sprintf("%d", 30+n)) }
func Bg(n int) string { return sgr(fmt.Sprintf("%d", 40+n)) }

func Fg256(n int) string { return sgr(fmt.Sprintf("38;5;%d", n)) }
func Bg256(n int) string { return sgr(fmt.Sprintf("48;5;%d", n)) }

func FgRGB(r, g, b int) string {
	return sgr(fmt.Sprintf("38;2;%d;%d;%d", r, g, b))
}
func BgRGB(r, g, b int) string {
	return sgr(fmt.Sprintf("48;2;%d;%d;%d", r, g, b))
}
