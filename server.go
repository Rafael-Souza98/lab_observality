package main

import (
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of requests handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"pong \n")
	pingCounter.Inc()
	
}

func main() {

	http.HandleFunc("/ping",ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8090", nil)
	
	prometheus.MustRegister(pingCounter)
}