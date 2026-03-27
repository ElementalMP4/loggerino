package style

func sgr(code string) string {
	return "\x1b[" + code + "m"
}

func Bold() string      { return sgr("1") }
func Dim() string       { return sgr("2") }
func Italic() string    { return sgr("3") }
func Underline() string { return sgr("4") }
func Blink() string     { return sgr("5") }
func Reverse() string   { return sgr("7") }
func Hidden() string    { return sgr("8") }
func Strike() string    { return sgr("9") }

func NoBold() string      { return sgr("22") }
func NoUnderline() string { return sgr("24") }
func NoBlink() string     { return sgr("25") }
func NoReverse() string   { return sgr("27") }
func Reset() string       { return sgr("0") }
