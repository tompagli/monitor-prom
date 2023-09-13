package local

import (
	"log"
	"html/template"

	api "github.com/hpe-usp-spire/signed-assertions/anonymousMode/api-libs/global"
	alOps "github.com/hpe-usp-spire/signed-assertions/anonymousMode/api-libs/options"
	"github.com/hpe-usp-spire/signed-assertions/anonymousMode/subject_workload/options"
<<<<<<< HEAD
=======
	"github.com/hpe-usp-spire/signed-assertions/anonymousMode/subject_workload/monitoring-prom"
>>>>>>> 117dace (prometheus instrumentation)
)

var Options *alOps.Options
var	Tpl		*template.Template

func init() {
<<<<<<< HEAD
=======
	monitor.RegisterMetric()
>>>>>>> 117dace (prometheus instrumentation)
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
