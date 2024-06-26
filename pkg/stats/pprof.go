package stats

import (
	"net/http"
	"net/http/pprof"

	deltapprof "github.com/grafana/pyroscope-go/godeltaprof/http/pprof"
)

func AddPprof(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/debug/pprof/delta_heap", deltapprof.Heap)
	mux.HandleFunc("/debug/pprof/delta_threadcreate", deltapprof.Mutex)
	mux.HandleFunc("/debug/pprof/delta_block", deltapprof.Block)

	profileDescriptions["/debug/pprof/"] = `PProf related things:<br/>
	<a href="/debug/pprof/goroutine?debug=2">full goroutine stack dump</a>
	`
}
