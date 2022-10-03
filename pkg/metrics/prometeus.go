package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var userStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_user_status_count", // metric name
		Help: "Count of status returned by user.",
	},
	[]string{"user", "status"}, // labels
)

func Register() error {
	return prometheus.Register(userStatus)
}
