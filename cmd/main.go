package main

import (
	"donow/src/config"
	"donow/src/logging"
	"donow/src/web/server"
	"os"
)

func cleanup() {
	log := logging.Get()
	exit_code := config.EXIT_SUCCESS_CODE
	if r := recover(); r != nil {
		log.Fatal(r.(string))
		exit_code = config.EXIT_ERROR_CODE
	}
	os.Exit(exit_code)
}

var (
	logfile string = "donow.log"
)

func main() {
	defer cleanup()

	log := logging.NewFile(logfile)
	log.Info("*** Starting Application ***")

	ipaddr, port := "0.0.0.0", 8086
	// ipaddr, port := "localhost", 8086
	err := server.Start(&ipaddr, port)
	if err != nil {
		log.Error(err.Error())
	}

	log.Info("*** Ending Application ***")
}
