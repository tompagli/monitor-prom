package main

import (
	// To sig. validation
	_ "crypto/sha256"

	"github.com/hpe-usp-spire/signed-assertions/anonymousMode/subject_workload/controller"
<<<<<<< HEAD
=======
	"github.com/hpe-usp-spire/signed-assertions/anonymousMode/subject_workload/monitoring-prom"
>>>>>>> 117dace (prometheus instrumentation)
	// Okta
)

func main() {
<<<<<<< HEAD
=======
	go monitor.PrometheusAPI()
	go monitor.UpdateMemoryUsage()
	go monitor.UpdateCPUUsage()
>>>>>>> 117dace (prometheus instrumentation)
	controller.SubjectWLController()
}
