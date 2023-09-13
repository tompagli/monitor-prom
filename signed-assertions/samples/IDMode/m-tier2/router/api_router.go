package router

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
<<<<<<< HEAD

=======
	muxprom "gitlab.com/msvechla/mux-prometheus/pkg/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
>>>>>>> 117dace (prometheus instrumentation)
	"github.com/hpe-usp-spire/signed-assertions/IDMode/m-tier2/handlers"
)

func MiddleTierRouter(ctx context.Context) (*mux.Router, error) {
	s := mux.NewRouter()

<<<<<<< HEAD
=======
	instrumentation := muxprom.NewDefaultInstrumentation()
	s.Use(instrumentation.Middleware)
	s.Path("/metrics").Handler(promhttp.Handler())

>>>>>>> 117dace (prometheus instrumentation)
	s.HandleFunc("/get_balance", handlers.GetBalanceHandler).Methods("POST")
	s.HandleFunc("/deposit", handlers.DepositHandler).Methods("POST")

	s.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	return s, nil
}
