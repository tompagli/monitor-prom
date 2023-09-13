package main

import (
	"context"

	"github.com/hpe-usp-spire/signed-assertions/IDMode/m-tier4/controller"
<<<<<<< HEAD
=======
	"github.com/hpe-usp-spire/signed-assertions/IDMode/m-tier4/monitoring-prom"
>>>>>>> 117dace (prometheus instrumentation)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
<<<<<<< HEAD
=======
	go monitor.PrometheusAPI()
	go monitor.UpdateMemoryUsage()
	go monitor.UpdateCPUUsage()
>>>>>>> 117dace (prometheus instrumentation)
	controller.MiddleTierController(ctx)
}
