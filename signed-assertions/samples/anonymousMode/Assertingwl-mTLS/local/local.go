package local

import (
	"log"

	"github.com/hpe-usp-spire/signed-assertions/anonymousMode/Assertingwl-mTLS/options"
	api "github.com/hpe-usp-spire/signed-assertions/anonymousMode/api-libs/global"
	alOps "github.com/hpe-usp-spire/signed-assertions/anonymousMode/api-libs/options"
	"github.com/hpe-usp-spire/signed-assertions/anonymousMode/Assertingwl-mTLS/monitoring-prom"
)

var Options *alOps.Options

func init() {
	monitor.RegisterRequestSizeMetric()
	log.Print("local init")
	// api-libs/options/options.go
	options, err := options.InitOptions()
	if err != nil {
		log.Fatal("Options init errored: ", err.Error())
	}

	Options = options
}

func InitGlobals() {
	log.Print("init global")

	// api-libs/global.go
	api.InitGlobals(Options)

}
