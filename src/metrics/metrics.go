package Metrics

type Metrics struct {
}

type MetriccImpl interface {
	Count(int)
	timer(float64)
}

func NewMetrics() *Metrics {
	return &Metrics{}
}
