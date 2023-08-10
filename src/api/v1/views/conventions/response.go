// Copyright [2023] [Martin Vuyk]
// https://github.com/martinvuyk/ginbartonic
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0

package conventions

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type ApiResponse[A any] struct {
	Success bool  `json:"success"`
	Data    *A    `json:"data"`
	Err     error `json:"err"`
	Code    uint  `json:"code"`
}
type endpointInterface[T, A any] interface {
	docs(c *gin.Context)
	handler(*gin.Context, T) (A, error)
}

type Endpoint[T, A any] struct {
	OkCode  int
	Url     string
	Metrics *Metrics
}

type HandlerType[T, A any] interface {
	func(*gin.Context, *T) (*A, error)
}

type ResponseType[T, A any, D HandlerType[T, A]] interface {
	func(c *gin.Context, in *T, handler D, code int) (*ApiResponse[A], error)
}

func Respond[T any, A any, D HandlerType[T, A]](c *gin.Context, in *T, handler D, code int, url string, metrics *Metrics) (*ApiResponse[A], error) {
	start := time.Now()
	resp, err := handler(c, in)
	if metrics != nil {
		monitorPrometheus(metrics, start, url, code, err)
	}
	if err != nil {
		return nil, err
	}
	return &ApiResponse[A]{Success: true, Code: uint(code), Data: resp}, nil
}

func monitorPrometheus(m *Metrics, start time.Time, url string, OkCode int, err error) {
	code := OkCode
	if err != nil {
		code = mapJujuError(err)
	}
	m.duration.With(prometheus.Labels{"endpoint": url, "status": fmt.Sprint(code)}).Observe(time.Since(start).Seconds())
}

type Metrics struct {
	duration *prometheus.HistogramVec
}

func NewMetrics(reg prometheus.Registerer, version string) *Metrics {
	m := &Metrics{
		duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: fmt.Sprintf("api_%s", version),
			Name:      "request_duration_seconds",
			Help:      "Duration of the request.",
			// 4 times larger for apdex score
			// Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
			// Buckets: prometheus.LinearBuckets(0.1, 5, 5),
			Buckets: []float64{0.00001, 0.00005, 0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.075,
				0.1, 0.15, 0.2, 0.25, 0.3, 0.5, 0.75, 0.9, 1.1, 1.5, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11, 12, 13, 14, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60},
		}, []string{"endpoint", "status"}),
	}
	reg.MustRegister(m.duration)
	return m
}
