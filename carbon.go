package carbon

import (
	"fmt"
	"net"
	"time"
)

type Carbon struct {
	Host     string
	Port     int
	Timeout  time.Duration
	conn     net.Conn
	Noop     bool
	Protocol string
}

// time in Seconds of how long each Carbon transaction is allowed to take
const defaultTimeout = 5

func (carbon *Carbon) IsNoop() bool {
	if carbon.Noop {
		return true
	}
	return false
}

// Is a Carbon Host defined? In other words, do we intend to send to carbon?
func (carbon *Carbon) IsDefined() bool {
	if carbon.Host == "" {
		return false
	}
	return true
}

func (carbon *Carbon) Connect() error {
	if carbon.IsDefined() {
		if carbon.conn != nil {
			carbon.conn.Close()
		}

		if !carbon.IsNoop() {
			address := fmt.Sprintf("%s:%d", carbon.Host, carbon.Port)
			conn, err := net.DialTimeout(carbon.Protocol, address, carbon.Timeout)
			if err != nil {
				return err
			}
			carbon.conn = conn
		}
	}
	return nil
}

func (carbon *Carbon) SendMetrics(metrics []Metric) error {
	if carbon.IsDefined() {
		if carbon.IsNoop() {
			return nil
		}
		for _, metric := range metrics {
			//			fmt.Printf("in SendMetrics, metric.String: %s\n", metric.String())
			_, err := fmt.Fprintf(carbon.conn, metric.String()+"\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Create new Carbon object as well as connect to Carbon instance.
func NewCarbon(host string, port int, noop bool) (*Carbon, error) {
	carbon := &Carbon{Host: host, Port: port, Timeout: time.Duration(1 * time.Minute), Protocol: "tcp", Noop: noop}
	err := carbon.Connect()
	if err != nil {
		return nil, err
	}
	return carbon, nil
}
