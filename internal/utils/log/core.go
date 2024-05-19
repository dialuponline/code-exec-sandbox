package log

/*
	log module is used to write log info to log file
	open a log file when log was created, and close it when log was destroyed
*/

import (
	"fmt"
	go_log "log"
	"os"
	"time"
)

type Log struct {
	Level int
	//File of log
	File *os.File
	path string
}

const (
	LOG_LEVEL_DEBUG = 0
	LOG_LEVEL_INFO  = 1
	LOG_LEVEL_WARN  = 2
	LOG_LEVEL_ERROR = 3
)

func (l *Log) Debug(format string, stdout bool, v ...interface{}) {
	if l.Level <= LOG_LEVEL_DEBUG {
		l.writeLog("DEBUG", format, stdout, v...)
	}
}

func (l *Log) Info(format string, stdout bool, v ...interface{}) {
	if l.Level <= LOG_LEVEL_INFO {
		l.writeLog("INFO", format, stdout, v...)
	}
}

func (l *Log) Warn(format string, stdout bool, v ...interface{}) {
	if l.Level <= LOG_LEVEL_WARN {
		l.writeLog("WARN", format, stdout, v...)
	}
}

func (l *Log) Error(format string, stdout bool, v ...interface{}) {
	if l.Level <= LOG_LEVEL_ERROR {
		l.writ