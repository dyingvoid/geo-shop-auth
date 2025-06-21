package metrics

type MetricClient interface {
	Inc(string)
	Update(string, float64)
}
