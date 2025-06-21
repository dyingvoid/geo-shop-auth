package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	RequestProcessTime  prometheus.Histogram
	RedisProcessTime    prometheus.Histogram
	PostgresProcessTime prometheus.Histogram
	RedisUp             prometheus.Gauge
	PostgresUp          prometheus.Gauge
}

func NewMetrics(namespace string) *Metrics {
	m := &Metrics{
		RequestProcessTime: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "request_process_duration_seconds",
			Help:      "Time taken to process http requests in seconds.",
			Buckets:   prometheus.DefBuckets,
		}),
		RedisProcessTime: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "redis_process_duration_seconds",
			Help:      "Time taken to process Redis queries in seconds.",
			Buckets:   prometheus.DefBuckets,
		}),
		PostgresProcessTime: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "postgres_process_duration_seconds",
			Help:      "Time taken to process Postgres queries in seconds.",
			Buckets:   prometheus.DefBuckets,
		}),
		RedisUp: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "redis_up",
			Help:      "Whether Redis is up (1) or down (0).",
		}),
		PostgresUp: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "postgres_up",
			Help:      "Whether Postgres is up (1) or down (0).",
		}),
	}

	prometheus.MustRegister(
		m.RequestProcessTime,
		m.RedisProcessTime,
		m.PostgresProcessTime,
		m.RedisUp,
		m.PostgresUp,
	)

	return m
}
