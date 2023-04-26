// Credits to
// https://antonputra.com/monitoring/monitor-golang-with-prometheus/

package prometheus

import (
	// "net/http"

	// "github.com/prometheus/client_golang/prometheus"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

// type metrics struct {
// 	devices prometheus.Gauge
// 	info    *prometheus.GaugeVec
// }

// func NewMetrics(reg prometheus.Registerer) *metrics {
// 	m := &metrics{
// 		devices: prometheus.NewGauge(prometheus.GaugeOpts{
// 			Namespace: "myapp",
// 			Name:      "connected_devices",
// 			Help:      "Number of currently connected devices.",
// 		}),
// 		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 			Namespace: "myapp",
// 			Name:      "info",
// 			Help:      "Information about the My App environment.",
// 		},
// 			[]string{"version"}),
// 	}
// 	reg.MustRegister(m.devices, m.info)
// 	return m
// }

// type Device struct {
// 	ID       int    `json:"id"`
// 	Mac      string `json:"mac"`
// 	Firmware string `json:"firmware"`
// }

// var dvs []Device

func Setup(router *gin.RouterGroup, reg *prometheus.Registry) {
	router.Any("/metrics", WrapHH(promhttp.Handler))

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	router.Any("/metrics/custom", WrapHH(func() http.Handler { return promHandler }))

	// reg := prometheus.NewRegistry()
	// m := NewMetrics(reg)

	// m.devices.Set(float64(len(dvs)))
	// version := "1.10.5"
	// m.info.With(prometheus.Labels{"version": version}).Set(1)

	// promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	// router.Any("/metrics/custom", WrapHH(func() http.Handler { return promHandler }))

}
