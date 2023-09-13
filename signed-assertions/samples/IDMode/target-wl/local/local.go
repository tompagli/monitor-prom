package local

import (
	"log"

	api "github.com/hpe-usp-spire/signed-assertions/IDMode/api-libs/global"
	alOps "github.com/hpe-usp-spire/signed-assertions/IDMode/api-libs/options"
	"github.com/hpe-usp-spire/signed-assertions/IDMode/target-wl/options"
<<<<<<< HEAD
=======
	"github.com/hpe-usp-spire/signed-assertions/IDMode/target-wl/monitoring-prom"
>>>>>>> 117dace (prometheus instrumentation)
)

var Options *alOps.Options

func init() {
<<<<<<< HEAD
=======
	monitor.RegisterRequestSizeMetric()
>>>>>>> 117dace (prometheus instrumentation)
	log.Print("local init")

	options, err := options.InitOptions()
	if err != nil {
		log.Fatal("Options init errored: ", err.Error())
	}

	Options = options
}

func InitGlobals() {
	log.Print("init local")
	api.InitGlobals(Options)

}

