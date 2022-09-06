package common

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe("0.0.0.0"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("Failed to start the monitor.")
		}
		log.Info("Monitor started successfully, port is: " + strconv.Itoa(port))
	}()
}
