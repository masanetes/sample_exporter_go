package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"runtime"
)

type collector struct {
	goroutinesDesc *prometheus.Desc
	threadsDesc    *prometheus.Desc
}

func newCollector() *collector {
	return &collector{
		goroutinesDesc: prometheus.NewDesc(
			"google_status",
			"Number of goroutines that currently exist.",
			nil, nil),
		threadsDesc: prometheus.NewDesc(
			"threads",
			"Number of OS threads created.",
			nil, nil),
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.goroutinesDesc
	ch <- c.threadsDesc
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {

	resp, _ := http.Get("https://google.com")

	log.Println("scrape")

	ch <- prometheus.MustNewConstMetric(c.goroutinesDesc, prometheus.GaugeValue, float64(resp.StatusCode))
	n, _ := runtime.ThreadCreateProfile(nil)
	ch <- prometheus.MustNewConstMetric(c.threadsDesc, prometheus.GaugeValue, float64(n))
}

func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(newCollector())

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	if err := http.ListenAndServe(":19115", nil); err != nil {
		log.Fatal(err)
	}
}
