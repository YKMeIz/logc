package logc

import "fmt"

const (
	White  = "\033[97m"
	Blue   = "\033[34m"
	Green  = "\033[97;32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
	Bold   = "\033[1m"
)

var (
	WhiteColor = func(v ...interface{}) string {
		return White + fmt.Sprint(v...) + Reset
	}

	BlueColor = func(v ...interface{}) string {
		return Blue + fmt.Sprint(v...) + Reset
	}

	GreenColor = func(v ...interface{}) string {
		return Green + fmt.Sprint(v...) + Reset
	}

	YellowColor = func(v ...interface{}) string {
		return Yellow + fmt.Sprint(v...) + Reset
	}

	RedColor = func(v ...interface{}) string {
		return Red + fmt.Sprint(v...) + Reset
	}
)
