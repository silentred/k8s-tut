package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"time"

	"crypto/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	id int32

	reqCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "req_count",
		Help: "total count of request",
	}, []string{"service"})

	reqDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "req_duration_us",
		Help: "latency of request in microsecond",
	}, []string{"service"})
)

func main() {
	binary.Read(rand.Reader, binary.LittleEndian, &id)

	// log
	logFile, err := os.Create("/tmp/hello-app.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	// prom
	prometheus.MustRegister(reqCount)
	prometheus.MustRegister(reqDuration)
	http.Handle("/metrics", prometheus.Handler())

	// http
	http.Handle("/", metricsMiddleware(helloWorld))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func metricsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		next(w, r)

		duration := float64(time.Now().Sub(t).Nanoseconds()) / 1000
		reqCount.WithLabelValues("http").Add(1)
		reqDuration.WithLabelValues("http").Observe(duration)
	})
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Printf("ID:%d time:%d", id, time.Now().Unix())
	fmt.Fprintf(w, "ID:%d path:%s", id, r.URL.Path)
}
