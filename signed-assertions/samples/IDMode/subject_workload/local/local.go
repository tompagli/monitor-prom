package local

import (
	"log"
	"html/template"

	api "github.com/hpe-usp-spire/signed-assertions/IDMode/api-libs/global"
	alOps "github.com/hpe-usp-spire/signed-assertions/IDMode/api-libs/options"
	"github.com/hpe-usp-spire/signed-assertions/IDMode/subject_workload/options"
	"github.com/hpe-usp-spire/signed-assertions/IDMode/subject_workload/monitoring-prom"
)

var Options *alOps.Options
var	Tpl		*template.Template

func init() {
	monitor.RegisterMetric()
	log.Print("local init")
	// api-libs/options/options.go
	options, err := options.InitOptions()
	if err != nil {
		log.Fatal("Options init errored: ", err.Error())
	}

	Options = options


}

func InitGlobals() {
	log.Print("init local")

	// api-libs/global.go
	api.InitGlobals(Options)

}

func InitTemplate() {
	Tpl 	= template.Must(template.ParseGlob("templates/*"))
}
