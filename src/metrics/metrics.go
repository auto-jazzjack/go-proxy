package Metrics

import (
	"time"

	collection "proxy/src/utils"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var buckets = [...]float64{0.5, 0.75, 0.9, 0.95, 0.99}

var emptyMap = map[string]string{}
var Label2xx = map[string]string{"status": "2xx"}
var Label3xx = map[string]string{"status": "3xx"}
var Label4xx = map[string]string{"status": "4xx"}
var Label5xx = map[string]string{"status": "5xx"}
var TOTAL_REQUEST_COUNT = "TOTAL_REQUEST_COUNT"
var TOTAL_REQUEST_LATENCY = "TOTAL_REQUEST_LATENCY"

func MeasureCountAndLatency(code int, from int64) {

	var clonedLabel = emptyMap
	if 200 <= code && code < 300 {
		clonedLabel = collection.CopyMap(Label2xx)
	} else if 300 <= code && code < 400 {
		clonedLabel = collection.CopyMap(Label3xx)
	} else if 400 <= code && code < 500 {
		clonedLabel = collection.CopyMap(Label4xx)
	} else if 500 <= code && code < 600 {
		clonedLabel = collection.CopyMap(Label5xx)
	} else {
		return
	}

	increaseTimer(clonedLabel, from)
	increaseCounter(clonedLabel)
}

func increaseCounter(label map[string]string) {
	newCounter(TOTAL_REQUEST_COUNT, label).Add(1)
}

func increaseTimer(label map[string]string, from int64) {

	now := time.Now().UnixMilli()
	elapsed := float64(now - from)

	newHistogram(TOTAL_REQUEST_LATENCY, label).Observe(elapsed)

}

func newCounter(name string, labels map[string]string) prometheus.Counter {
	return promauto.NewCounter(
		prometheus.CounterOpts{
			Name:        name,
			ConstLabels: labels,
		},
	)
}

func newHistogram(name string, labels map[string]string) prometheus.Histogram {

	return promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:        name,
			ConstLabels: labels,
			Buckets:     buckets[:],
		},
	)
}
