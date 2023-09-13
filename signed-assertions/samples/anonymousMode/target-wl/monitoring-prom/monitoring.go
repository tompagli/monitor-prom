package monitor

import (
    "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
    "net/http"
    "runtime"
    "time"
    "github.com/shirou/gopsutil/cpu"
)

var (
    CpuUsage = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "app_cpu_usage_percentage",
            Help: "Current CPU usage percentage of the application.",
        },
    )
)

var (
    MemoryUsage = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "app_memory_usage_bytes",
            Help: "Current memory usage of the application in bytes.",
        },
    )
)

func UpdateCPUUsage() {
    for {
        percentages, err := cpu.Percent(1*time.Second, false)
        if err == nil && len(percentages) > 0 {
            // Update the cpuUsage metric with the current CPU usage percentage
            CpuUsage.Set(percentages[0])
        }

        // Sleep for a specific interval (e.g., 1 minute)
        time.Sleep(10 * time.Millisecond)
    }
}

func UpdateMemoryUsage() {
    for {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)

        // Update the memoryUsage metric with the current memory usage in bytes
        MemoryUsage.Set(float64(m.Alloc))

        // Sleep for a specific interval (e.g., 1 minute)
        time.Sleep(10 * time.Millisecond)
    }
}


func RegisterRequestSizeMetric() {
    defaultRegistry := prometheus.NewRegistry()
    prometheus.DefaultRegisterer = defaultRegistry
    prometheus.DefaultGatherer = defaultRegistry
    prometheus.MustRegister(MemoryUsage)
    prometheus.MustRegister(CpuUsage)
}

func PrometheusAPI(){
	http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":2114", nil)
}


