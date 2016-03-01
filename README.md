carbon-golang
=============

# What is carbon-golang?

A Go-based library written to make it easier to send Carbon metrics to a Carbon cache.

# How do I use it?

From your binary, 

```golang
import (
    "github.com/jforman/carbon-golang"
)

package main

func main() {
    var string carbonHost
    var int carbonPort
    var bool carbonNoop
    var bool carbonVerbose
    carbonReceiver, err := carbon.NewCarbon(carbonHost, carbonPort, carbonNoop, carbonVerbose)
    metrics := []carbon.Metric
    metrics := append(metrics, carbon.Metric{Name: "foo.bar.min", Value: 2, Timestamp: 1234567890})
    metrics := append(metrics, carbon.Metric{Name: "foo.bar.max", Value: 10, Timestamp: 1234567890})
    carbonReciever.SendMetrics(metrics)
}
```

Or in shot: Create an instance of a carbon receiver, make a list of metrics, and send the metrics in.

# What's next?

I plan to refactor the code to make it more correct, and add helper functions to make sending metrics in even more straightforward and easy.
