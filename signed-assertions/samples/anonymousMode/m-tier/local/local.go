package local

import (
	"log"

	api "github.com/hpe-usp-spire/signed-assertions/anonymousMode/api-libs/global"
	alOps "github.com/hpe-usp-spire/signed-assertions/anonymousMode/api-libs/options"
	"github.com/hpe-usp-spire/signed-assertions/anonymousMode/m-tier/options"
	"github.com/hpe-usp-spire/signed-assertions/anonymousMode/m-tier/monitoring-prom"
)

var Options *alOps.Options

func init() {
	monitor.RegisterRequestSizeMetric()
	log.Print("global init")
	
	options, err := options.InitOptions()
	if err != nil {
		log.Fatal("Options init errored: ", err.Error())
	}

	Options = options
}

func InitGlobals() {
	log.Print("init global")
	api.InitGlobals(Options)
}
