package carbon

import (
	"fmt"
	//	"strconv"
)

type Metric struct {
	Name      string
	Value     float64
	Timestamp int64
}

func (metric Metric) String() string {
	//timestampAsString, _ := strconv.Itoa(metric.Timestamp)
	return fmt.Sprintf("%s %f %d", metric.Name,
		metric.Value, metric.Timestamp)
}
