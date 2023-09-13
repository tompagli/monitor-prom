package main

import (
	"context"

	"github.com/hpe-usp-spire/signed-assertions/IDMode/m-tier4/controller"
	"github.com/hpe-usp-spire/signed-assertions/IDMode/m-tier4/monitoring-prom"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go monitor.PrometheusAPI()
	go monitor.UpdateMemoryUsage()
	go monitor.UpdateCPUUsage()
	controller.MiddleTierController(ctx)
}
