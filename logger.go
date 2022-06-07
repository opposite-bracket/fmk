package fmk

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

const txtKey = "txtId"

type logLevel int

const (
	INFO    logLevel = 1
	DEBUG            = 2
	WARNING          = 3
	ERROR            = 4
)

// apiLog holds apiLogger information such as
//		* timetamp
//		* logLevel (INFO, DEBUG, WARNING, ERROR)
//      * txId (optional)
type apiLog struct {
	logLevel  logLevel
	common    map[string]interface{}
	startTime time.Time
	endTime   time.Time
}

// apiLogger is the apiLogger singleton
var apiLogger *apiLog

// once makes the logger concurrent-safe
var once sync.Once

// LogLevel sets the level of the apiLogger
// 		INFO    logLevel = 1
//		DEBUG            = 2
//		WARNING          = 3
//		ERROR            = 4
func (l *apiLog) LogLevel(level logLevel) {
	l.logLevel = level
}

func (l *apiLog) StartTimer() {
	l.startTime = time.Now()
}

func (l *apiLog) EndTimer() {
	l.endTime = time.Now()
}

func (l *apiLog) ResetTxId() {
	l.common[txtKey] = uuid.NewString()
}

func (l *apiLog) Logf(format string, v ...any) {
	levelStr := ""
	switch l.logLevel {
	case INFO:
		levelStr = "INFO"
	case DEBUG:
		levelStr = "DEBUG"
	case WARNING:
		levelStr = "WARNING"
	case ERROR:
		levelStr = "ERROR"
	}

	context := map[string]interface{}{
		"payload":   fmt.Sprintf(format, v...),
		"level":     levelStr,
		"timestamp": time.Now().Format("2006-01-02 15:04:05.000000"),
	}

	for k, v := range l.common {
		context[k] = v
	}

	data, _ := json.Marshal(context)
	fmt.Println(string(data))
}

func (l *apiLog) Log(v ...any) {
	levelStr := ""
	switch l.logLevel {
	case INFO:
		levelStr = "INFO"
	case DEBUG:
		levelStr = "DEBUG"
	case WARNING:
		levelStr = "WARNING"
	case ERROR:
		levelStr = "ERROR"
	}

	context := map[string]interface{}{
		"payload":   fmt.Sprint(v...),
		"level":     levelStr,
		"timestamp": time.Now().Format("2006-01-02 15:04:05.000000"),
	}

	for k, v := range l.common {
		context[k] = v
	}

	data, _ := json.Marshal(context)
	fmt.Println(string(data))
}

// ApiLog gets a concurrent-safe api apiLogger instance
func ApiLog() *apiLog {
	once.Do(func() {
		apiLogger = &apiLog{
			logLevel: INFO,
			common:   make(map[string]interface{}, 0),
		}
	})

	return apiLogger
}
