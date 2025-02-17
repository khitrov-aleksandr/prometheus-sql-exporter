package main

import (
	"log"
	"monitoring/prometheus-sql-exporter/config"
	"monitoring/prometheus-sql-exporter/db"
	"monitoring/prometheus-sql-exporter/metric"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	countCalls prometheus.Gauge
	cpuTemp    prometheus.Gauge
	hdFailures *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		countCalls: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "count_all_calls",
			Help: "Count call for all projects",
		}),
		cpuTemp: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_temperature_celsius",
			Help: "Current temperature of the CPU.",
		}),
		hdFailures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "hd_errors_total",
				Help: "Number of hard-disk errors.",
			},
			[]string{"device"},
		),
	}
	reg.MustRegister(m.countCalls)
	reg.MustRegister(m.cpuTemp)
	reg.MustRegister(m.hdFailures)
	return m
}

func main() {
	cfg := config.New()
	db := db.New(cfg)

	metric := metric.NewMetric(db)

	// Create a non-global registry.
	reg := prometheus.NewRegistry()

	// Create new metrics and register them using the custom registry.
	m := NewMetrics(reg)

	//Set count calls
	m.countCalls.Set(float64(metric.GetCountCall()))

	// Set values for the new created metrics.
	m.cpuTemp.Set(65.3)
	m.hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	// Expose metrics and custom registry via an HTTP server
	// using the HandleFor function. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Fatal(http.ListenAndServe(":9200", nil))
}
