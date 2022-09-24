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
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("启动失败")
		}
		log.Info("监控启动,端口为：" + strconv.Itoa(port))
	}()
}
