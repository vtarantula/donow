package logging

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

type msgType uint

const (
	INFO msgType = iota + 0
	DEBUG
	ERROR
	FATAL
	WARNING
)

const (
	STDOUT string = "stdout"
	STDERR string = "stderr"
)

type logConfig struct {
	logfd    *os.File
	filename string
}

var (
	l_fileLogger  []*logConfig
	defaultLogger *logConfig
)

func init() {
	stdout_logger := &logConfig{
		logfd:    os.Stdout,
		filename: STDOUT,
	}
	stderr_logger := &logConfig{
		logfd:    os.Stderr,
		filename: STDERR,
	}

	defaultLogger = stdout_logger
	l_fileLogger = append(l_fileLogger, stdout_logger)
	l_fileLogger = append(l_fileLogger, stderr_logger)
}

func writeToFile(fd *os.File, message_type msgType, message *string) {

	mt := &sync.Mutex{}
	mt.Lock()

	var message_prefix string
	// Default case is not necessary because message type is a enum
	switch message_type {
	case INFO:
		message_prefix = "INFO"
	case DEBUG:
		message_prefix = "DEBUG"
	case ERROR:
		message_prefix = "ERROR"
	case WARNING:
		message_prefix = "WARNING"
	case FATAL:
		// Not using logger.Fatal to take power of application crashing away from logging module
		message_prefix = "FATAL"
	}

	log.SetOutput(fd)
	logger := log.New(fd, fmt.Sprintf("%-8s: ", message_prefix), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	logger.Printf("%s", *message)
	mt.Unlock()
}

func (o_log *logConfig) Info(message string) {
	writeToFile(o_log.logfd, INFO, &message)
}

func (o_log *logConfig) Warning(message string) {
	writeToFile(o_log.logfd, WARNING, &message)
}

func (o_log *logConfig) Error(message string) {
	var fd *os.File = o_log.logfd
	if o_log.logfd == os.Stdout {
		fd = os.Stderr
	}
	writeToFile(fd, ERROR, &message)
}

func (o_log *logConfig) Debug(message string) {
	writeToFile(o_log.logfd, DEBUG, &message)
}

func (o_log *logConfig) Fatal(message string) {
	var fd *os.File = o_log.logfd
	if o_log.logfd == os.Stdout {
		fd = os.Stderr
	}
	writeToFile(fd, FATAL, &message)
}

// NewFile returns a logger for using the filename passed
func NewFile(filename string) *logConfig {
	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range l_fileLogger {
		if v.filename == filename {
			return v
		}
	}

	logger := &logConfig{
		logfd:    fd,
		filename: fd.Name(),
	}

	defaultLogger = logger

	l_fileLogger = append(l_fileLogger, logger)

	return logger
}

// New returns a logger for using stdout
func New() *logConfig {
	var logger *logConfig
	for _, v := range l_fileLogger {
		if v.logfd == os.Stdout {
			logger = v
			break
		}
	}
	return logger
}

// New returns the default logger
// By default, it will point to the latest file logConfig object
// If no file has been configured for logging, it will return stdout logConfig
func Get() *logConfig {
	return defaultLogger
}

// Set returns the default logger
func Set(filename *string) error {
	for _, v := range l_fileLogger {
		if v.filename == *filename {
			defaultLogger = v
			return nil
		}
	}
	return errors.New("unable to find existing logger")
}

// Cleanup cleans up all the file descriptors
func Cleanup() {
	for _, v := range l_fileLogger {
		v.logfd.Close()
	}
}
