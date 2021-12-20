package Metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var buckets = [...]float64{0.5, 0.75, 0.9, 0.95, 0.99}

var emptyMap = map[string]string{}

var LABEL_MAP = map[int]string{2: "2xx", 3: "3xx", 4: "4xx", 5: "5xx"}
var TOTAL_REQUEST_COUNT = "TOTAL_REQUEST_COUNT"
var TOTAL_REQUEST_LATENCY = "TOTAL_REQUEST_LATENCY"
var METRIC_LABEL = []string{"status", "path"}

var cachedMetricObject = make(map[string]GeneralMetric)

type GeneralMetric struct {
	counter   *prometheus.CounterVec
	histogram *prometheus.HistogramVec
}

func MeasureCountAndLatency(code int, uri string, from int64) {

	var clonedLabel = []string{}

	clonedLabel = append(clonedLabel, LABEL_MAP[code/100])
	clonedLabel = append(clonedLabel, uri)

	if _, ok := cachedMetricObject[uri]; !ok {
		cachedMetricObject[uri] = GeneralMetric{
			counter:   newCounter(TOTAL_REQUEST_COUNT),
			histogram: newHistogram(TOTAL_REQUEST_LATENCY),
		}
	}

	if len(clonedLabel) != 2 {
		return
	}

	now := time.Now().UnixMilli()
	elapsed := float64(now - from)

	cachedMetricObject[uri].counter.WithLabelValues(clonedLabel...).Add(1)
	cachedMetricObject[uri].histogram.WithLabelValues(clonedLabel...).Observe(elapsed)

}

func newCounter(name string) *prometheus.CounterVec {
	return promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
		},
		METRIC_LABEL,
	)
}

func newHistogram(name string) *prometheus.HistogramVec {

	return promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    name,
			Buckets: buckets[:],
		},
		METRIC_LABEL,
	)
}
