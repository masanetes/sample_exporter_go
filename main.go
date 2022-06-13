package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type collector struct {
	goroutinesDesc *prometheus.Desc
	target         string
}

func newCollector(target string) *collector {
	return &collector{
		goroutinesDesc: prometheus.NewDesc(
			"http_status",
			"http status",
			nil, nil),
		target: target,
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.goroutinesDesc
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	resp, err := http.Get(c.target)
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.goroutinesDesc, prometheus.GaugeValue, 0)
	} else {
		ch <- prometheus.MustNewConstMetric(c.goroutinesDesc, prometheus.GaugeValue, float64(resp.StatusCode))
	}
}

func probeHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	reg := prometheus.NewRegistry()
	reg.MustRegister(newCollector(target))
	h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/probe", probeHandler)
	if err := http.ListenAndServe(":19115", nil); err != nil {
		log.Fatal(err)
	}
}
