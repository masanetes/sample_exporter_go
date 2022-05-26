package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func recordMetrics() {
	for {
		opsProcessed.Inc()
		time.Sleep(2 * time.Second)
	}
}

func main() {
	go recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":19115", nil); err != nil {
		log.Fatal(err)
	}
}
