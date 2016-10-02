package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/op/go-logging"
	"os"
)

var logfile_path = "imgserve.log"
var log = logging.MustGetLogger("imgserve")

var stdout_format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var file_format = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} - %{level:.4s} %{id:03x} %{message}`,
)

var Conf mainConfig

func main() {
	// Setup logging
	logStdout := logging.NewLogBackend(os.Stdout, "", 0)
	logStdoutFormatter := logging.NewBackendFormatter(logStdout, stdout_format)

	// Logging to file
	logfh, err := os.OpenFile(logfile_path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Could not open log file %s", logfile_path)
		panic(err)
	}
	defer logfh.Close()
	logFile := logging.NewLogBackend(logfh, "", 0)
	logFileFormatter := logging.NewBackendFormatter(logFile, file_format)
	/* To limit logging to a level
	logFileLeveled := logging.AddModuleLevel(logFile)
	logFileLeveled.SetLevel(logging.ERROR, "")
	*/

	// Start logging
	logging.SetBackend(logStdoutFormatter, logFileFormatter)

	log.Debug("Starting up...")
	Conf = ReadConfig()
	log.Debug(Conf.Path)
	api := iris.New()
	// Register handlers
	for path, handlerfunc := range GetHandlerMap() {
		api.Get(path, handlerfunc)
	}
	for path, handlerfunc := range PostHandlerMap() {
		api.Post(path, handlerfunc)
	}
	api.Listen(":8080")
}
