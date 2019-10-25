package logc

import (
	"fmt"
	"reflect"
)

var (
	TimeFormat    = "[2006-01-02 15:04:05]"
	DefaultFormat = func(v ...interface{}) string {
		return defaultFormat(WhiteColor, v...)
	}
	DebugFormat = func(v ...interface{}) string {
		return defaultFormat(BlueColor, v...)
	}
	ImportantFormat = func(v ...interface{}) string {
		return defaultFormat(GreenColor, v...)
	}
	WarningFormat = func(v ...interface{}) string {
		return defaultFormat(YellowColor, v...)
	}
	ErrorFormat = func(v ...interface{}) string {
		return defaultFormat(RedColor, v...)
	}
)

func defaultFormat(color func(...interface{}) string, v ...interface{}) string {
	var text = fmt.Sprint(v...)

	if !ColorMode {
		switch reflect.ValueOf(color) {
		case reflect.ValueOf(BlueColor):
			text = "[DEBUG] " + text
		case reflect.ValueOf(GreenColor):
			text = "[IMPORTANT] " + text
		case reflect.ValueOf(YellowColor):
			text = "[WARNING] " + text
		case reflect.ValueOf(RedColor):
			text = "[ERROR] " + text
		default:
		}
	}

	if TimeStamp {
		text = GetTime() + " " + text
	}

	if ColorMode {
		if ColorInBold {
			text = Bold + text
		}
		return color(text)
	}

	return text
}
