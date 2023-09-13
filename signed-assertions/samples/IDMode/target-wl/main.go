package main

import (
	"context"

	"github.com/hpe-usp-spire/signed-assertions/IDMode/target-wl/controller"
<<<<<<< HEAD
)

func main() {
=======
	"github.com/hpe-usp-spire/signed-assertions/IDMode/target-wl/monitoring-prom"
)

func main() {
	go monitor.PrometheusAPI()
	go monitor.UpdateMemoryUsage()
	go monitor.UpdateCPUUsage()
>>>>>>> 117dace (prometheus instrumentation)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	controller.TargetWLController(ctx)
}
