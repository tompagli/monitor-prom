package main

import (
	"context"

	"github.com/hpe-usp-spire/signed-assertions/anonymousMode/m-tier5/controller"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go monitor.PrometheusAPI()
	go monitor.UpdateMemoryUsage()
	go monitor.UpdateCPUUsage()
	controller.MiddleTierController(ctx)
}
