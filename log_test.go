package logc

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestDefault(t *testing.T) {
	t.Log("default configuration.")
	helper(t, Default, "a test message in default format.", White+GetTime()+" a test message in default format."+Reset)

	t.Log("log level of 6.")
	LogLevel = 6
	helper(t, Default, "a test message in default format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Default, "a test message in default format.", White+"a test message in default format."+Reset)
	TimeStamp = true
}

func TestDefaultBold(t *testing.T) {
	ColorInBold = true

	t.Log("default configuration with bold font.")
	helper(t, Default, "a test message in default format.", White+Bold+GetTime()+" a test message in default format."+Reset)

	t.Log("log level of 6.")
	LogLevel = 6
	helper(t, Default, "a test message in default format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Default, "a test message in default format.", White+Bold+"a test message in default format."+Reset)
	TimeStamp = true

	ColorInBold = false
}

func TestDefaultNoColor(t *testing.T) {
	ColorMode = false

	t.Log("default configuration with no color.")
	helper(t, Default, "a test message in default format.", GetTime()+" a test message in default format.")

	t.Log("log level of 6.")
	LogLevel = 6
	helper(t, Default, "a test message in default format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Default, "a test message in default format.", "a test message in default format.")
	TimeStamp = true

	ColorMode = true
}

func TestDebug(t *testing.T) {
	t.Log("default configuration.")
	helper(t, Debug, "a test message in debug format.", Blue+GetTime()+" a test message in debug format."+Reset)

	t.Log("log level of 8.")
	LogLevel = 8
	helper(t, Debug, "a test message in debug format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Debug, "a test message in debug format.", Blue+"a test message in debug format."+Reset)
	TimeStamp = true
}

func TestDebugBold(t *testing.T) {
	ColorInBold = true

	t.Log("default configuration with bold font.")
	helper(t, Debug, "a test message in debug format.", Blue+Bold+GetTime()+" a test message in debug format."+Reset)

	t.Log("log level of 8.")
	LogLevel = 8
	helper(t, Debug, "a test message in debug format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Debug, "a test message in debug format.", Blue+Bold+"a test message in debug format."+Reset)
	TimeStamp = true

	ColorInBold = false
}

func TestDebugNoColor(t *testing.T) {
	ColorMode = false

	t.Log("default configuration with no color.")
	helper(t, Debug, "a test message in debug format.", GetTime()+" [DEBUG] a test message in debug format.")

	t.Log("log level of 8.")
	LogLevel = 8
	helper(t, Debug, "a test message in debug format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Debug, "a test message in debug format.", "[DEBUG] a test message in debug format.")
	TimeStamp = true

	ColorMode = true
}

func TestImportant(t *testing.T) {
	t.Log("default configuration.")
	helper(t, Important, "a test message in important format.", Green+GetTime()+" a test message in important format."+Reset)

	t.Log("log level of 4.")
	LogLevel = 4
	helper(t, Important, "a test message in important format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Important, "a test message in important format.", Green+"a test message in important format."+Reset)
	TimeStamp = true
}

func TestImportantBold(t *testing.T) {
	ColorInBold = true

	t.Log("default configuration with bold font.")
	helper(t, Important, "a test message in important format.", Green+Bold+GetTime()+" a test message in important format."+Reset)

	t.Log("log level of 4.")
	LogLevel = 4
	helper(t, Important, "a test message in important format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Important, "a test message in important format.", Green+Bold+"a test message in important format."+Reset)
	TimeStamp = true

	ColorInBold = false
}

func TestImportantNoColor(t *testing.T) {
	ColorMode = false

	t.Log("default configuration with no color.")
	helper(t, Important, "a test message in important format.", GetTime()+" [IMPORTANT] a test message in important format.")

	t.Log("log level of 4.")
	LogLevel = 4
	helper(t, Important, "a test message in important format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Important, "a test message in important format.", "[IMPORTANT] a test message in important format.")
	TimeStamp = true

	ColorMode = true
}

func TestWarning(t *testing.T) {
	t.Log("default configuration.")
	helper(t, Warning, "a test message in warning format.", Yellow+GetTime()+" a test message in warning format."+Reset)

	t.Log("log level of 2.")
	LogLevel = 2
	helper(t, Warning, "a test message in warning format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Warning, "a test message in warning format.", Yellow+"a test message in warning format."+Reset)
	TimeStamp = true
}

func TestWarningBold(t *testing.T) {
	ColorInBold = true

	t.Log("default configuration with bold font.")
	helper(t, Warning, "a test message in warning format.", Yellow+Bold+GetTime()+" a test message in warning format."+Reset)

	t.Log("log level of 2.")
	LogLevel = 2
	helper(t, Warning, "a test message in warning format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Warning, "a test message in warning format.", Yellow+Bold+"a test message in warning format."+Reset)
	TimeStamp = true

	ColorInBold = false
}

func TestWarningNoColor(t *testing.T) {
	ColorMode = false

	t.Log("default configuration with no color.")
	helper(t, Warning, "a test message in warning format.", GetTime()+" [WARNING] a test message in warning format.")

	t.Log("log level of 2.")
	LogLevel = 2
	helper(t, Warning, "a test message in warning format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Warning, "a test message in warning format.", "[WARNING] a test message in warning format.")
	TimeStamp = true

	ColorMode = true
}

func TestError(t *testing.T) {
	ErrorWithExit = false

	t.Log("default configuration.")
	helper(t, Error, "a test message in error format.", Red+GetTime()+" a test message in error format."+Reset)

	t.Log("log level of 0.")
	LogLevel = 0
	helper(t, Error, "a test message in error format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Error, "a test message in error format.", Red+"a test message in error format."+Reset)
	TimeStamp = true
}

func TestErrorBold(t *testing.T) {
	ErrorWithExit = false
	ColorInBold = true

	t.Log("default configuration with bold font.")
	helper(t, Error, "a test message in error format.", Red+Bold+GetTime()+" a test message in error format."+Reset)

	t.Log("log level of 0.")
	LogLevel = 0
	helper(t, Error, "a test message in error format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Error, "a test message in error format.", Red+Bold+"a test message in error format."+Reset)
	TimeStamp = true

	ColorInBold = false
}

func TestErrorNoColor(t *testing.T) {
	ErrorWithExit = false
	ColorMode = false

	t.Log("default configuration with no color.")
	helper(t, Error, "a test message in error format.", GetTime()+" [ERROR] a test message in error format.")

	t.Log("log level of 0.")
	LogLevel = 0
	helper(t, Error, "a test message in error format.", "")
	LogLevel = 9

	t.Log("non-timestamp mode.")
	TimeStamp = false
	helper(t, Error, "a test message in error format.", "[ERROR] a test message in error format.")
	TimeStamp = true

	ColorMode = true
}

// helper tests the function.
// It accepts testing type T, function needs to be tested, input value, and expected output value.
func helper(t *testing.T, f interface{}, in, e string) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	}()

	f.(func(...interface{}))(in)

	bufString := buf.String()

	if len(bufString) > 1 {
		bufString = bufString[:len(buf.String())-1]
	}

	if strings.Compare(bufString, e) != 0 {
		funcName := filepath.Base(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		t.Error(funcName, "return wrong output:\n expect:\n", e, "\n get output:\n", bufString)
		return
	}

	t.Log(bufString)
}
