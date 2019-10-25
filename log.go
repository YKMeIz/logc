package logc

import (
	"log"
	"os"
	"runtime/debug"
	"time"
)

func init() {
	log.SetFlags(0)
}

func GetTime() string {
	return time.Now().Format(TimeFormat)
}

func Debug(v ...interface{}) {
	if LogLevel >= 9 {
		log.Println(DebugFormat(v...))
	}
}

func Default(v ...interface{}) {
	if LogLevel >= 7 {
		log.Println(DefaultFormat(v...))
	}
}

func Important(v ...interface{}) {
	if LogLevel >= 5 {
		log.Println(ImportantFormat(v...))
	}
}

func Warning(v ...interface{}) {
	if LogLevel >= 3 {
		log.Println(WarningFormat(v...))
	}
}

func Error(v ...interface{}) {
	if LogLevel >= 1 {
		if ErrorWithPanic {
			log.Println(ErrorFormat(string(debug.Stack())))
			os.Exit(1)
		}
		if ErrorWithExit {
			log.Fatal(ErrorFormat(v...))
		}
		log.Println(ErrorFormat(v...))
	}
	if ErrorWithExit {
		os.Exit(1)
	}
}
