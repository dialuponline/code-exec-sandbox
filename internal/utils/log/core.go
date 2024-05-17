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
	path 