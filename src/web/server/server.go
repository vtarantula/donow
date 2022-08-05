package server

import (
	"donow/src/logging"
	"errors"
	"fmt"
	"net"
	"net/http"
)

func Start(ipaddr *string, port int) error {

	log := logging.Get()

	if *ipaddr != "localhost" {
		ip := net.ParseIP(*ipaddr)
		if ip == nil {
			log.Error("Not a valid IP: " + *ipaddr)
			return errors.New("unable to validate IP " + *ipaddr)
		}
	}

	address := fmt.Sprintf("%s:%d", *ipaddr, port)
	log.Info("Starting webserver using address: " + address)

	mux := http.NewServeMux()

	setRoutes(mux)

	http.ListenAndServe(address, mux)

	return nil
}
