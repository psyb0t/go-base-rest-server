package main

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func ReadyLogTypes(logfile *os.File) {
	Trace = log.New(io.MultiWriter(logfile, os.Stdout),
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(logfile, os.Stdout),
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(io.MultiWriter(logfile, os.Stdout),
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(logfile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func SetupLogging() {
	var err error

	var logfile *os.File
	logfile, err = os.OpenFile(logfile_path,
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		logfile = os.Stderr
	}

	ReadyLogTypes(logfile)
}
